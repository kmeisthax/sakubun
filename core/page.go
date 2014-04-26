package core

import "html/template"

type ResourceTag int

const (
    StyleResource ResourceTag = iota
    ScriptResource
)

/* Identifies an external resource to be imported into a Page.
 */
type Resource struct {
    /* Identifies what kind of resource is to be imported.
     * 
     * A StyleResource is a Resource that imports styling information to the
     * page (typically, text/css). Styling information affects the content of
     * the page and is imported using a rel="stylesheet" link.
     * 
     * A ScriptResource is a Resource that imports new program code into the
     * client browsing context (typically, application/javascript). Scripts
     * add additional behaviors to the page and are imported using a script
     * element.
     */
    Type ResourceTag
    
    /* Specify the exact mime type of the resource.
     * 
     * If left blank, the default MIME type for this resource type will be used
     * (text/css for styles and application/javascript for scripts).
     * 
     * If not left blank, this will override the specified MIME type. This can
     * be used to indicate an alternate scripting or styling language to be
     * consumed by a client-side resource compiler or compatible browser.
     */
    MimeType string
    
    /* Public URL of the resource. Ideallly should be a domain-relative URL.
     */
    Url template.URL
}

type Header struct {
    Title string
    Resources []Resource
}

type Page struct {
    Language string
    Header Header
}