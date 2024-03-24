package schemas

type KeplerApiCaseCommonParam struct {
	Key       string `json:"key"`
	Value     string `json:"value"`
	ValueType string `json:"valueType"`
	Required  bool   `json:"required"`
	Enabled   bool   `json:"enabled"`
	Desc      string `json:"desc"`
}

type KeplerApiCaseHeaders struct {
	KeplerApiCaseCommonParam
}

type KeplerApiCaseParams struct {
	KeplerApiCaseCommonParam
}

type KeplerApiCaseRest struct {
	RId       string `json:"rId"`
	Key       string `json:"key"`
	Name      string `json:"name"`
	Required  bool   `json:"required"`
	ValueType string `json:"valueType"`
	Value     string `json:"value"`
	Desc      string `json:"desc"`
}

type KeplerApiCaseOutputParams struct {
	RId       string `json:"rId"`
	Name      string `json:"name"`
	Key       string `json:"key"`
	ValueType string `json:"valueType"`
	Value     string `json:"value"`
	Desc      string `json:"desc"`
}

type KeplerApiCaseValidations struct {
	RId             string `json:"rId"`
	ValueType       string `json:"valueType"`
	Name            string `json:"name"`
	Key             string `json:"key"`
	Value           string `json:"value"`
	DisableKeyInput bool   `json:"disableKeyInput"`
}

type KeplerApiCase struct {
	Id            int                         `json:"id"`
	Headers       []KeplerApiCaseHeaders      `json:"headers"`
	Params        []KeplerApiCaseParams       `json:"params"`
	Rest          []KeplerApiCaseRest         `json:"rest"`
	OutputParams  []KeplerApiCaseOutputParams `json:"outputParams"`
	Validations   []KeplerApiCaseValidations  `json:"validations"`
	HttpType      string                      `json:"httpType"`
	Method        string                      `json:"method"`
	ModifierName  string                      `json:"modifier_name"`
	CreatorName   string                      `json:"creator_name"`
	VersionName   *string                     `json:"version_name"`
	PreNode       interface{}                 `json:"pre_node"`
	PostNode      interface{}                 `json:"post_node"`
	ApiTags       []string                    `json:"api_tags"`
	Mark          string                      `json:"mark"`
	Priority      string                      `json:"priority"`
	Description   string                      `json:"description"`
	Host          string                      `json:"host"`
	Path          string                      `json:"path"`
	Body          string                      `json:"body"`
	BodyType      string                      `json:"bodyType"`
	Checkcode     int                         `json:"checkcode"`
	Checkvalue    string                      `json:"checkvalue"`
	Checkkey      string                      `json:"checkkey"`
	Timeout       int                         `json:"timeout"`
	Retry         int                         `json:"retry"`
	PreCode       string                      `json:"pre_code"`
	PostCode      string                      `json:"post_code"`
	Sign          string                      `json:"sign"`
	Ak            string                      `json:"ak"`
	Ctime         string                      `json:"ctime"`
	Mtime         string                      `json:"mtime"`
	IsDeleted     bool                        `json:"is_deleted"`
	RetryInterval int                         `json:"retry_interval"`
	IsLocked      bool                        `json:"is_locked"`
	Yapi          int                         `json:"yapi"`
	Creator       string                      `json:"creator"`
	Modifier      string                      `json:"modifier"`
	Version       *string                     `json:"version"`
}
