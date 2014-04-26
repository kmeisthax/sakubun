package core

import (
    "html/template"
    "path/filepath"
    "reflect"
    "io"
    "string"
)

/* A theme is a collection of templates and resources to be used together.
 */
type Theme struct {
    /* The name of this theme. Must match the name of the template in the
     * ThemeRegistry.
     */
    Name string

    /* The name of a base theme which can be queried for additional templates.
     * Must exist within the theme registry or be a blank string.
     */
    BaseTheme string

    /* The directory where the theme is stored.
     * The directory is scanned for *.tpl files which are parsed and stored in
     * the Templates list.
     */
    Dir string

    /* List of templates provided by the theme (not including base themes).
     * The templates stored here should not be associated with any context.
     */
    Templates map[string]*template.Template

    /* Master template which holds all templates applicable to this theme,
     * including base theme templates.
     */
    TmplCtxt *template.Template

    hasBeenScanned bool
}

/* The theme registry is the list of all known themes and templates.
 */
var ThemeRegistry map[string]*Theme

/* Overlays template maps such that the final template map contains all
 * templates present within all maps.
 *
 * If a template is present in multiple maps, the map specified first wins.
 */
func overlayTemplateMap(tlist ...map[string]*template.Template) map[string]*template.Template {
    var out map[string]*template.Template

    for _, tMap := range tlist {
        for tName, tPtr := range tMap {
            if out[tName] == nil {
                out[tName] = tPtr
            }
        }
    }

    return out
}

/* Register a theme
 */
func RegisterTheme(th *Theme) error {
    matches, ok := filepath.Glob(filepath.Join(th.Dir, "*.tpl"))
    if (ok != nil) {
        return ok
    }

    for _, matchedPath := range matches {
        newTemplate, ok := template.ParseFiles(matchedPath)
        if (ok != nil) {
            return ok
        }

        th.Templates[strings.ToLower(newTemplate.Name())] = newTemplate
    }

    tplBase := ThemeRegistry[th.BaseTheme]
    tplOverlay := overlayTemplateMap(th.Templates)
    for tplBase != nil {
        tplOverlay = overlayTemplateMap(tplOverlay, tplBase.Templates)
        tplBase = ThemeRegistry[tplBase.BaseTheme]
    }

    th.TmplCtxt = template.New("__master__")

    for tName, tPtr := range tplOverlay {
        th.TmplCtxt.AddParseTree(tName, tPtr.Tree)
    }

    th.hasBeenScanned = true
    return nil
}

/* Given assembled data, render it with an applicable theme.
 * 
 * Data is expected to be a struct or map of string to interface{}.
 * If the data is of a type with a name, then that name will be used for the
 * template. Otherwise, the data is expected to have a Theme field or key
 * naming what template is to be used.
 */
func RenderData(th *Theme, data *interface{}, writer io.Writer) error {
    data_type := reflect.TypeOf(data)
    if data_type.Kind() == reflect.Ptr {
        data_type = data_type.Elem()
    }
    
    data_value := reflect.ValueOf(data)
    //Variables we expect to extract from the data structure
    theme_name := strings.ToLower(data_type.Name())
    
    if data_type.Kind() == reflect.Struct {
        theme_field, ok := data_type.FieldByName("Theme")
        if !ok {
            return nil
        }
        
        if theme_field.Type.Kind() == reflect.String {
            theme_value := data_value.FieldByName("Theme")
            theme_name = theme_value.String()
        }
    } else {
        return nil
    }
    
    return th.TmplCtxt.ExecuteTemplate(writer, theme_name, data)
}