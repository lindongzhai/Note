package schemas

type KeplerApiDocProperties struct {
	Format    string `json:"format"`
	Default   string `json:"default"`
	Maximum   string `json:"maximum"`
	Minimum   string `json:"minimum"`
	MaxLength string `json:"maxLength"`
	MinLength string `json:"minLength"`
	Pattern   string `json:"pattern"`
	Enum      string `json:"enum"`
	Example   string `json:"example"`
}

type KeplerApiDocCommonParam struct {
	Key        string                 `json:"key"`
	Value      []interface{}          `json:"value"`
	ValueType  string                 `json:"valueType"`
	Required   bool                   `json:"required"`
	Properties KeplerApiDocProperties `json:"properties"`
	Desc       string                 `json:"desc"`
}

type KeplerApiDocHeaders struct {
	KeplerApiDocCommonParam
}

type KeplerApiDocParam struct {
	KeplerApiDocCommonParam
	In string `json:"in"`
}

type KeplerApiDocCommonBodyData struct {
	KeplerApiDocCommonParam
	RefValue interface{} `json:"refValue"`
	RefName  interface{} `json:"refName"`
}

type KeplerApiDocBodyJsonData struct {
	KeplerApiDocCommonBodyData
}

type KeplerApiDocBodyJsonDataArray struct {
	KeplerApiDocCommonBodyData
}

type KeplerApiDocBodyFormData struct {
	KeplerApiDocCommonBodyData
}

type KeplerApiDocBodyXWwwFormUrlencoded struct {
	KeplerApiDocCommonBodyData
}

type KeplerApiDocBody struct {
	None               interface{}                          `json:"none"`
	JSON               string                               `json:"JSON"`
	Text               string                               `json:"Text"`
	JavaScript         string                               `json:"JavaScript"`
	HTML               string                               `json:"HTML"`
	XML                string                               `json:"XML"`
	JsonData           []KeplerApiDocBodyJsonData           `json:"json-data"`
	JsonDataArray      []KeplerApiDocBodyJsonDataArray      `json:"json-data-array"`
	FormData           []KeplerApiDocBodyFormData           `json:"form-data"`
	XWwwFormUrlencoded []KeplerApiDocBodyXWwwFormUrlencoded `json:"x-www-form-urlencoded"`
	Binary             string                               `json:"binary"`
}

type KeplerApiDocResponse struct {
	Headers     []KeplerApiDocHeaders `json:"headers"`
	Body        KeplerApiDocBody      `json:"body"`
	BodyType    string                `json:"bodyType"`
	StatusCode  int                   `json:"status_code"`
	Description string                `json:"description"`
	Id          int                   `json:"id"`
}

type KeplerApiDoc struct {
	Id            int                    `json:"id"`
	Headers       KeplerApiDocHeaders    `json:"headers"`
	Params        KeplerApiDocParam      `json:"params"`
	Body          KeplerApiDocBody       `json:"body"`
	OutputParams  []interface{}          `json:"outputParams"`
	Validations   []interface{}          `json:"validations"`
	Protocol      []string               `json:"protocol"`
	Method        string                 `json:"method"`
	Responses     []KeplerApiDocResponse `json:"responses"`
	YApiTags      []string               `json:"yapi_tags"`
	Manager       string                 `json:"manager"`
	ModuleNames   []string               `json:"module_names"`
	Cases         []KeplerApiCase        `json:"cases"`
	Description   string                 `json:"description"`
	Remarks       string                 `json:"remarks"`
	Status        string                 `json:"status"`
	Host          string                 `json:"host"`
	Path          string                 `json:"path"`
	Sign          string                 `json:"sign"`
	Ak            string                 `json:"ak"`
	BodyType      string                 `json:"bodyType"`
	Timeout       int                    `json:"timeout"`
	Retry         int                    `json:"retry"`
	PreCode       string                 `json:"pre_code"`
	PostCode      string                 `json:"post_code"`
	IsDeleted     bool                   `json:"is_deleted"`
	RetryInterval int                    `json:"retry_interval"`
	IsLocked      bool                   `json:"is_locked"`
	IsPatrolled   bool                   `json:"is_patrolled"`
	Priority      string                 `json:"priority"`
}
