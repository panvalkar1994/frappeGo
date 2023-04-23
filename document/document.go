package document

import (
	"encoding/json"
	"fmt"
	"go/format"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

const timeLayout = "2006-01-02 15:04:05.000000"

var skipFields = [...]string{
	"Section Break",
	"Button",
	"Column Break",
	"Fold",
	"Geolocation",
}

var fieldMap = map[string]string{
	"Attach":       "string",
	"Attach Image": "string",
	"Barcode":      "string",
	"Check":        "uint",
	"Code":         "string",
	"Color":        "string",
	"Currency":     "float64",
	"Data":         "string",
	"Date":         "time.Time",
	"Datetime":     "time.Time",
	"Dynamic Link": "string",
	"Float":        "float64",
	"Heading":      "string",
	"HTML":         "string",
	"HTML Editor":  "string",
	"Image":        "string",
	"Int":          "int64",
	"Link":         "string",
	"Long Text":    "string",
	"Password":     "string",
	"Percent":      "float64",
	"Read Only":    "uint",
	"Select":       "string",
	"Small Text":   "string",
	"Table":        "string",
	"Text":         "string",
	"Text Editor":  "string",
	"Time":         "time.Time",
	"Signature":    "string",
}

type DocField struct {
	Name                  string `json:"name"`
	Creation              string `json:"creation"`
	Modified              string `json:"modified"`
	ModifiedBy            string `json:"modified_by"`
	Owner                 string `json:"owner"`
	Docstatus             uint   `json:"docstatus"`
	Parent                string `json:"parent"`
	Parentfield           string `json:"parentfield"`
	Parenttype            string `json:"parenttype"`
	Idx                   uint   `json:"idx"`
	Fieldname             string `json:"fieldname"`
	Label                 string `json:"label"`
	Oldfieldname          string `json:"oldfieldname"`
	Fieldtype             string `json:"fieldtype"`
	Oldfieldtype          string `json:"oldfieldtype"`
	Options               string `json:"options"`
	SearchIndex           uint   `json:"search_index"`
	Hidden                uint   `json:"hidden"`
	SetOnlyOnce           uint   `json:"set_only_once"`
	AllowInQuickEntry     uint   `json:"allow_in_quick_entry"`
	PrintHide             uint   `json:"print_hide"`
	ReportHide            uint   `json:"report_hide"`
	Reqd                  uint   `json:"reqd"`
	Bold                  uint   `json:"bold"`
	InGlobalSearch        uint   `json:"in_global_search"`
	Collapsible           uint   `json:"collapsible"`
	Unique                uint   `json:"unique"`
	NoCopy                uint   `json:"no_copy"`
	AllowOnSubmit         uint   `json:"allow_on_submit"`
	Trigger               string `json:"trigger"`
	CollapsibleDependsOn  string `json:"collapsible_depends_on"`
	DependsOn             string `json:"depends_on"`
	Permlevel             uint   `json:"permlevel"`
	IgnoreUserPermissions uint   `json:"ignore_user_permissions"`
	Width                 string `json:"width"`
	PrintWidth            string `json:"print_width"`
	Columns               uint   `json:"columns"`
	Default               string `json:"default"`
	Description           string `json:"description"`
	InListView            uint   `json:"in_list_view"`
	InStandardFilter      uint   `json:"in_standard_filter"`
	ReadOnly              uint   `json:"read_only"`
	Precision             string `json:"precision"`
	Length                uint   `json:"length"`
	Translatable          uint   `json:"translatable"`
}

type DocType struct {
	Name                   string `json:"name"`
	Creation               string `json:"creation"`
	Modified               string `json:"modified"`
	ModifiedBy             string `json:"modified_by"`
	Owner                  string `json:"owner"`
	Docstatus              uint   `json:"docstatus"`
	Parent                 string `json:"parent"`
	Parentfield            string `json:"parentfield"`
	Parenttype             string `json:"parenttype"`
	Idx                    uint   `json:"idx"`
	SearchFields           string `json:"search_fields"`
	IsSingle               uint   `json:"issingle"`
	IsTable                uint   `json:"istable"`
	EditableGrid           uint   `json:"editable_grid"`
	TrackChanges           uint   `json:"track_changes"`
	Module                 string `json:"module"`
	RestrictToDomain       string `json:"restrict_to_domain"`
	App                    string `json:"app"`
	Autoname               string `json:"autoname"`
	NameCase               string `json:"name_case"`
	TitleField             string `json:"title_field"`
	ImageField             string `json:"image_field"`
	TimelineField          string `json:"timeline_field"`
	SortField              string `json:"sort_field"`
	SortOrder              string `json:"sort_order"`
	Description            string `json:"description"`
	Colour                 string `json:"colour"`
	ReadOnly               uint   `json:"read_only"`
	InCreate               uint   `json:"in_create"`
	MenuIndex              uint   `json:"menu_index"`
	ParentNode             string `json:"parent_node"`
	SmallIcon              string `json:"smallicon"`
	AllowCopy              uint   `json:"allow_copy"`
	AllowRename            uint   `json:"allow_rename"`
	AllowImport            uint   `json:"allow_import"`
	HideToolbar            uint   `json:"hide_toolbar"`
	HideHeading            uint   `json:"hide_heading"`
	TrackSeen              uint   `json:"track_seen"`
	MaxAttachments         uint   `json:"max_attachments"`
	PrintOutline           string `json:"print_outline"`
	ReadOnlyOnload         uint   `json:"read_only_onload"`
	DocumentType           string `json:"document_type"`
	Icon                   string `json:"icon"`
	Color                  string `json:"color"`
	TagFields              string `json:"tag_fields"`
	Subject                string `json:"subject"`
	LastUpdate             string `json:"_last_update"`
	Engine                 string `json:"engine"`
	DefaultPrintFormat     string `json:"default_print_format"`
	IsSubmittable          uint   `json:"is_submittable"`
	ShowNameInGlobalSearch uint   `json:"show_name_in_global_search"`
	UserTags               string `json:"_user_tags"`
	Custom                 uint   `json:"custom"`
	Beta                   uint   `json:"beta"`
	ImageView              uint   `json:"image_view"`
	HasWebView             uint   `json:"has_web_view"`
	AllowGuestToView       uint   `json:"allow_guest_to_view"`
	Route                  string `json:"route"`
	IsPublishedField       string `json:"is_published_field"`
}

type DocPerm struct {
	Name                  string `json:"name"`
	Creation              string `json:"creation"`
	Modified              string `json:"modified"`
	ModifiedBy            string `json:"modified_by"`
	Owner                 string `json:"owner"`
	Docstatus             uint   `json:"docstatus"`
	Parent                string `json:"parent"`
	Parentfield           string `json:"parentfield"`
	Parenttype            string `json:"parenttype"`
	Idx                   uint   `json:"idx"`
	Fieldname             string `json:"fieldname"`
	Label                 string `json:"label"`
	Oldfieldname          string `json:"oldfieldname"`
	Fieldtype             string `json:"fieldtype"`
	Oldfieldtype          string `json:"oldfieldtype"`
	Options               string `json:"options"`
	Search_index          uint   `json:"search_index"`
	Hidden                uint   `json:"hidden"`
	SetOnlyOnce           uint   `json:"set_only_once"`
	AllowInQuickEntry     uint   `json:"allow_in_quick_entry"`
	PrintHide             uint   `json:"print_hide"`
	ReportHide            uint   `json:"report_hide"`
	Reqd                  uint   `json:"reqd"`
	Bold                  uint   `json:"bold"`
	InGlobalSearch        uint   `json:"in_global_search"`
	Collapsible           uint   `json:"collapsible"`
	Unique                uint   `json:"unique"`
	NoCopy                uint   `json:"no_copy"`
	AllowOnSubmit         uint   `json:"allow_on_submit"`
	Trigger               string `json:"trigger"`
	CollapsibleDependsOn  string `json:"collapsible_depends_on"`
	DependsOn             string `json:"depends_on"`
	Permlevel             uint   `json:"permlevel"`
	IgnoreUserPermissions uint   `json:"ignore_user_permissions"`
	Width                 string `json:"width"`
	PrintWidth            string `json:"print_width"`
	Columns               uint   `json:"columns"`
	Default               string `json:"default"`
	Description           string `json:"description"`
	InListView            uint   `json:"in_list_view"`
	InStandardFilter      uint   `json:"in_standard_filter"`
	ReadOnly              uint   `json:"read_only"`
	Precision             string `json:"precision"`
	Length                uint   `json:"length"`
	Translatable          uint   `json:"translatable"`
}

type Document struct {
	DocType
	Fields      []DocField `json:"fields"`
	Permissions []DocPerm  `json:"permissions"`
}

func (d *DocType) GetModified() (time.Time, error) {
	t, err := time.Parse(timeLayout, d.Modified)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

func ParseDocument(filePath string) (Document, error) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return Document{}, err
	}

	var document Document
	err = json.Unmarshal(data, &document)
	if err != nil {
		return Document{}, err
	}
	return document, nil
}

func (d *Document) CreateDocumentModel(dest string) {

	if dest == "" {
		dest = "./models"
	}

	err := os.MkdirAll(dest, os.ModePerm)

	filePath := fmt.Sprintf("%s/%s.go", dest, scrubName(d.Name))
	_, err = os.Stat(filePath)
	var file *os.File
	if err == nil {
		file, err = os.Open(filePath)
		if err != nil {
			fmt.Println("error 1: ", err)
		}
	} else {
		var err1 error
		file, err1 = os.Create(filePath)
		if err1 != nil {
			fmt.Println("error 2: ", err1)
		}
	}
	file.Write(d.createDocSource())
}

func (d *Document) createDocSource() []byte {
	out := fmt.Sprintf("type %s struct {\n", d.Name)
	for _, v := range d.Fields {
		if skipThisField(v.Fieldtype) {
			continue
		}
		out += fmt.Sprintf("\t%s\t%s\t`json:\"%s\"`\n", scrubName(v.Label), fieldMap[v.Fieldtype], v.Fieldname)
	}
	out += "}\n"
	s, e := format.Source([]byte(out))
	if e != nil {
		fmt.Println("error 3: ", 3)
	}
	return s
}

func skipThisField(Fieldtype string) bool {
	for _, skipField := range skipFields {
		if skipField == Fieldtype {
			return true
		}
	}
	return false
}

func scrubName(field string) string {
	words := strings.Split(field, " ")
	var out string
	for _, word := range words {
		out += word
	}
	return out
}
