package core

import (
    "html/template"
    "path/filepath"
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

        th.Templates[newTemplate.Name()] = newTemplate
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
