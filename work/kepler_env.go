package schemas

type KeplerEnvExtraHeaders struct {
	GenericHeader int    `json:"generic_header"`
	ReplaceMode   int    `json:"replace_mode"`
	HeaderName    string `json:"header_name"`
}

type KeplerEnvExtraHeadersConfig struct {
	GenericHeader string `json:"generic_header"`
	ReplaceMode   int    `json:"replace_mode"`
}

type KeplerEnvExtra struct {
	Host          string                      `json:"host"`
	Scheme        string                      `json:"scheme"`
	Headers       KeplerEnvExtraHeaders       `json:"headers"`
	Enabled       bool                        `json:"enabled"`
	HeadersConfig KeplerEnvExtraHeadersConfig `json:"headers_config"`
}

type KeplerEnvTeamEnvs struct {
	Id           int         `json:"id"`
	TeamId       int         `json:"teamid"`
	Name         string      `json:"name"`
	Value        string      `json:"value"`
	Description  string      `json:"description"`
	EnvName      string      `json:"env_name"`
	Env          string      `json:"env"`
	ValueType    string      `json:"valueType"`
	Creator      string      `json:"creator"`
	Modifier     string      `json:"modifier"`
	CreatorName  string      `json:"creator_name"`
	ModifierName string      `json:"modifier_name"`
	Ctime        string      `json:"ctime"`
	Mtime        string      `json:"mtime"`
	AccountId    interface{} `json:"account_id"`
	AccountName  interface{} `json:"account_name"`
}

type KeplerEnv struct {
	Id                int                 `json:"id"`
	Extra             KeplerEnvExtra      `json:"extra"`
	Name              string              `json:"name"`
	Description       string              `json:"description"`
	Env               string              `json:"env"`
	SpEnv             string              `json:"sp_env"`
	IsPrivate         bool                `json:"is_private"`
	EnabledPermission bool                `json:"enabled_permission"`
	Scope             int                 `json:"scope"`
	TeamId            int                 `json:"teamid"`
	Creator           string              `json:"creator"`
	Modifier          string              `json:"modifier"`
	ProjectId         interface{}         `json:"project_id"`
	ProjectHosts      []interface{}       `json:"projecthosts"`
	TeamEnvs          []KeplerEnvTeamEnvs `json:"teamenvs"`
}
