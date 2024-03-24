package schemas

type KeplerScenarioCaseStepExtendsCondition struct {
	Expect string `json:"expect"`
	Op     string `json:"op"`
	Actual string `json:"actual"`
}

type KeplerScenarioCaseStepExtendsLoop struct {
	Expect    string `json:"expect,omitempty"`
	Op        string `json:"op,omitempty"`
	Actual    string `json:"actual,omitempty"`
	Interval  int    `json:"interval,omitempty"`
	Timeout   int    `json:"timeout,omitempty"`
	IterValue string `json:"iter_value,omitempty"`
	Iterable  string `json:"iterable,omitempty"`
}

type KeplerScenarioCaseStepExtends struct {
	Step      int                                     `json:"step"`
	KwArgs    interface{}                             `json:"kwargs"`
	RefMulApi interface{}                             `json:"ref_mulapi"`
	Condition *KeplerScenarioCaseStepExtendsCondition `json:"condition"`
	Loop      *KeplerScenarioCaseStepExtendsLoop      `json:"loop"`
}

type KeplerScenarioCaseStep struct {
	Id         int                            `json:"id"`
	Api        *KeplerApiCase                 `json:"api"`
	WsApi      interface{}                    `json:"wsapi"`
	Extends    *KeplerScenarioCaseStepExtends `json:"extends"`
	Desc       string                         `json:"desc"`
	Type       string                         `json:"type"`
	Delay      int                            `json:"delay"`
	Priority   int                            `json:"priority"`
	Enabled    bool                           `json:"enabled"`
	Ctime      string                         `json:"ctime"`
	Mtime      string                         `json:"mtime"`
	IsExecuted bool                           `json:"is_executed"`
	MulApi     int                            `json:"mulApi"`
	Parent     *interface{}                   `json:"parent"`
}

type KeplerScenarioCase struct {
	Id           int                      `json:"id"`
	ModifierName string                   `json:"modifier_name"`
	CreatorName  string                   `json:"creator_name"`
	ServiceName  string                   `json:"service_name"`
	VersionName  interface{}              `json:"version_name"`
	Steps        []KeplerScenarioCaseStep `json:"steps"`
	Script       *string                  `json:"script"`
	ApiDataS     []interface{}            `json:"apidatas"`
	ModuleNames  []string                 `json:"module_names"`
	TestType     string                   `json:"testType"`
	Description  string                   `json:"description"`
	Autogenerate bool                     `json:"autogenerate"`
	Ctime        string                   `json:"ctime"`
	Mtime        string                   `json:"mtime"`
	Priority     string                   `json:"priority"`
	Plugin       string                   `json:"plugin"`
	Tag          string                   `json:"tag"`
	Model        string                   `json:"model"`
	IsDeleted    bool                     `json:"is_deleted"`
	IsLocked     bool                     `json:"is_locked"`
	IsTerminated bool                     `json:"is_terminated"`
	TeamId       int                      `json:"teamid"`
	ProjectId    int                      `json:"projectid"`
	ServiceId    int                      `json:"serviceid"`
	Creator      string                   `json:"creator"`
	Modifier     string                   `json:"modifier"`
	Version      *interface{}             `json:"version"`
	User         string                   `json:"user"`
	Maker        string                   `json:"maker"`
}
