package elasticsearch

import "github.com/chi-chu/log/define"

type esMapping struct {
	Settings		*settings				`json:"settings"`
	Mappings		mappings				`json:"mappings"`
}

type mappings struct {
	Properties		map[string]*Field		`json:"properties"`
}

type settings struct {
	Shards			int						`json:"number_of_shards"`
	Replicas		int						`json:"number_of_replicas"`
}

type Field struct {
	Type			string					`json:"type"`
	Index			bool					`json:"index,omitempty"`
	Format			string					`json:"format,omitempty"`
}

type EsConfig struct {
	Shards			int
	Replicas		int
	FieldMap		map[string]*Field
	bulkInsert		bool
	sniff			bool
}

var defaultFieldMap = map[string]*Field{
	define.TIPS_FILE:	&Field{Type:"text"},
	define.TIPS_LEVEL:	&Field{Type:"text"},
	define.TIPS_FUNC:	&Field{Type:"text"},
	define.TIPS_LINE:	&Field{Type:"text"},
	define.TIPS_MSG:	&Field{Type:"keyword"},
	define.TIPS_TIME:	&Field{Type:"keyword", Format:"yyyy-MM-dd HH:mm:ss"},
}

var defaultConfig = &EsConfig{
	Shards:   1,
	Replicas: 2,
	FieldMap: defaultFieldMap,
	bulkInsert: false,
	sniff:		false,
}

type Option func(*EsConfig)

func opt(option ...Option) {
	for _,f := range option {
		f(defaultConfig)
	}
}

func SetEsConfig(config *EsConfig) Option {
	return func(cf *EsConfig) {
		defaultConfig = cf
	}
}

func SetShards(num int) Option {
	return func(cf *EsConfig) {
		cf.Shards = num
	}
}

func SetReplicas(num int) Option {
	return func(cf *EsConfig) {
		cf.Replicas = num
	}
}

func SetBulk(b bool) Option {
	return func(cf *EsConfig) {
		cf.bulkInsert = b
	}
}

func SetAddField(column string, field *Field) Option {
	return func(cf *EsConfig) {
		if field != nil {
			cf.FieldMap[column] = field
		}
	}
}

func SetSniff(b bool) Option {
	return func(cf *EsConfig) {
		cf.sniff = b
	}
}

func newMap() *esMapping {
	return &esMapping{
		Settings: newSetting(),
		Mappings: mappings{Properties:defaultFieldMap},
	}
}

func newSetting() *settings {
	return &settings{defaultConfig.Shards,defaultConfig.Replicas}
}