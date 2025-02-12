// Generated by godel-refreshable-plugin: do not edit.

package httpclient

import refreshable "github.com/palantir/pkg/refreshable"

type RefreshableServicesConfig interface {
	refreshable.Refreshable
	CurrentServicesConfig() ServicesConfig
	MapServicesConfig(func(ServicesConfig) interface{}) refreshable.Refreshable
	SubscribeToServicesConfig(func(ServicesConfig)) (unsubscribe func())

	Default() RefreshableClientConfig
	Services() RefreshableStringToClientConfig
}

type RefreshingServicesConfig struct {
	refreshable.Refreshable
}

func NewRefreshingServicesConfig(in refreshable.Refreshable) RefreshingServicesConfig {
	return RefreshingServicesConfig{Refreshable: in}
}

func (r RefreshingServicesConfig) CurrentServicesConfig() ServicesConfig {
	return r.Current().(ServicesConfig)
}

func (r RefreshingServicesConfig) MapServicesConfig(mapFn func(ServicesConfig) interface{}) refreshable.Refreshable {
	return r.Map(func(i interface{}) interface{} {
		return mapFn(i.(ServicesConfig))
	})
}

func (r RefreshingServicesConfig) SubscribeToServicesConfig(consumer func(ServicesConfig)) (unsubscribe func()) {
	return r.Subscribe(func(i interface{}) {
		consumer(i.(ServicesConfig))
	})
}

func (r RefreshingServicesConfig) Default() RefreshableClientConfig {
	return NewRefreshingClientConfig(r.MapServicesConfig(func(i ServicesConfig) interface{} {
		return i.Default
	}))
}

func (r RefreshingServicesConfig) Services() RefreshableStringToClientConfig {
	return NewRefreshingStringToClientConfig(r.MapServicesConfig(func(i ServicesConfig) interface{} {
		return i.Services
	}))
}

type RefreshableClientConfig interface {
	refreshable.Refreshable
	CurrentClientConfig() ClientConfig
	MapClientConfig(func(ClientConfig) interface{}) refreshable.Refreshable
	SubscribeToClientConfig(func(ClientConfig)) (unsubscribe func())

	ServiceName() refreshable.String
	URIs() refreshable.StringSlice
	APIToken() refreshable.StringPtr
	APITokenFile() refreshable.StringPtr
	DisableHTTP2() refreshable.BoolPtr
	ProxyFromEnvironment() refreshable.BoolPtr
	ProxyURL() refreshable.StringPtr
	MaxNumRetries() refreshable.IntPtr
	InitialBackoff() refreshable.DurationPtr
	MaxBackoff() refreshable.DurationPtr
	ConnectTimeout() refreshable.DurationPtr
	ReadTimeout() refreshable.DurationPtr
	WriteTimeout() refreshable.DurationPtr
	IdleConnTimeout() refreshable.DurationPtr
	TLSHandshakeTimeout() refreshable.DurationPtr
	ExpectContinueTimeout() refreshable.DurationPtr
	HTTP2ReadIdleTimeout() refreshable.DurationPtr
	HTTP2PingTimeout() refreshable.DurationPtr
	MaxIdleConns() refreshable.IntPtr
	MaxIdleConnsPerHost() refreshable.IntPtr
	Metrics() RefreshableMetricsConfig
	Security() RefreshableSecurityConfig
}

type RefreshingClientConfig struct {
	refreshable.Refreshable
}

func NewRefreshingClientConfig(in refreshable.Refreshable) RefreshingClientConfig {
	return RefreshingClientConfig{Refreshable: in}
}

func (r RefreshingClientConfig) CurrentClientConfig() ClientConfig {
	return r.Current().(ClientConfig)
}

func (r RefreshingClientConfig) MapClientConfig(mapFn func(ClientConfig) interface{}) refreshable.Refreshable {
	return r.Map(func(i interface{}) interface{} {
		return mapFn(i.(ClientConfig))
	})
}

func (r RefreshingClientConfig) SubscribeToClientConfig(consumer func(ClientConfig)) (unsubscribe func()) {
	return r.Subscribe(func(i interface{}) {
		consumer(i.(ClientConfig))
	})
}

func (r RefreshingClientConfig) ServiceName() refreshable.String {
	return refreshable.NewString(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.ServiceName
	}))
}

func (r RefreshingClientConfig) URIs() refreshable.StringSlice {
	return refreshable.NewStringSlice(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.URIs
	}))
}

func (r RefreshingClientConfig) APIToken() refreshable.StringPtr {
	return refreshable.NewStringPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.APIToken
	}))
}

func (r RefreshingClientConfig) APITokenFile() refreshable.StringPtr {
	return refreshable.NewStringPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.APITokenFile
	}))
}

func (r RefreshingClientConfig) DisableHTTP2() refreshable.BoolPtr {
	return refreshable.NewBoolPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.DisableHTTP2
	}))
}

func (r RefreshingClientConfig) ProxyFromEnvironment() refreshable.BoolPtr {
	return refreshable.NewBoolPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.ProxyFromEnvironment
	}))
}

func (r RefreshingClientConfig) ProxyURL() refreshable.StringPtr {
	return refreshable.NewStringPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.ProxyURL
	}))
}

func (r RefreshingClientConfig) MaxNumRetries() refreshable.IntPtr {
	return refreshable.NewIntPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.MaxNumRetries
	}))
}

func (r RefreshingClientConfig) InitialBackoff() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.InitialBackoff
	}))
}

func (r RefreshingClientConfig) MaxBackoff() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.MaxBackoff
	}))
}

func (r RefreshingClientConfig) ConnectTimeout() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.ConnectTimeout
	}))
}

func (r RefreshingClientConfig) ReadTimeout() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.ReadTimeout
	}))
}

func (r RefreshingClientConfig) WriteTimeout() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.WriteTimeout
	}))
}

func (r RefreshingClientConfig) IdleConnTimeout() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.IdleConnTimeout
	}))
}

func (r RefreshingClientConfig) TLSHandshakeTimeout() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.TLSHandshakeTimeout
	}))
}

func (r RefreshingClientConfig) ExpectContinueTimeout() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.ExpectContinueTimeout
	}))
}

func (r RefreshingClientConfig) HTTP2ReadIdleTimeout() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.HTTP2ReadIdleTimeout
	}))
}

func (r RefreshingClientConfig) HTTP2PingTimeout() refreshable.DurationPtr {
	return refreshable.NewDurationPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.HTTP2PingTimeout
	}))
}

func (r RefreshingClientConfig) MaxIdleConns() refreshable.IntPtr {
	return refreshable.NewIntPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.MaxIdleConns
	}))
}

func (r RefreshingClientConfig) MaxIdleConnsPerHost() refreshable.IntPtr {
	return refreshable.NewIntPtr(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.MaxIdleConnsPerHost
	}))
}

func (r RefreshingClientConfig) Metrics() RefreshableMetricsConfig {
	return NewRefreshingMetricsConfig(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.Metrics
	}))
}

func (r RefreshingClientConfig) Security() RefreshableSecurityConfig {
	return NewRefreshingSecurityConfig(r.MapClientConfig(func(i ClientConfig) interface{} {
		return i.Security
	}))
}

type RefreshableMetricsConfig interface {
	refreshable.Refreshable
	CurrentMetricsConfig() MetricsConfig
	MapMetricsConfig(func(MetricsConfig) interface{}) refreshable.Refreshable
	SubscribeToMetricsConfig(func(MetricsConfig)) (unsubscribe func())

	Enabled() refreshable.BoolPtr
	Tags() RefreshableStringToString
}

type RefreshingMetricsConfig struct {
	refreshable.Refreshable
}

func NewRefreshingMetricsConfig(in refreshable.Refreshable) RefreshingMetricsConfig {
	return RefreshingMetricsConfig{Refreshable: in}
}

func (r RefreshingMetricsConfig) CurrentMetricsConfig() MetricsConfig {
	return r.Current().(MetricsConfig)
}

func (r RefreshingMetricsConfig) MapMetricsConfig(mapFn func(MetricsConfig) interface{}) refreshable.Refreshable {
	return r.Map(func(i interface{}) interface{} {
		return mapFn(i.(MetricsConfig))
	})
}

func (r RefreshingMetricsConfig) SubscribeToMetricsConfig(consumer func(MetricsConfig)) (unsubscribe func()) {
	return r.Subscribe(func(i interface{}) {
		consumer(i.(MetricsConfig))
	})
}

func (r RefreshingMetricsConfig) Enabled() refreshable.BoolPtr {
	return refreshable.NewBoolPtr(r.MapMetricsConfig(func(i MetricsConfig) interface{} {
		return i.Enabled
	}))
}

func (r RefreshingMetricsConfig) Tags() RefreshableStringToString {
	return NewRefreshingStringToString(r.MapMetricsConfig(func(i MetricsConfig) interface{} {
		return i.Tags
	}))
}

type RefreshableStringToString interface {
	refreshable.Refreshable
	CurrentStringToString() map[string]string
	MapStringToString(func(map[string]string) interface{}) refreshable.Refreshable
	SubscribeToStringToString(func(map[string]string)) (unsubscribe func())
}

type RefreshingStringToString struct {
	refreshable.Refreshable
}

func NewRefreshingStringToString(in refreshable.Refreshable) RefreshingStringToString {
	return RefreshingStringToString{Refreshable: in}
}

func (r RefreshingStringToString) CurrentStringToString() map[string]string {
	return r.Current().(map[string]string)
}

func (r RefreshingStringToString) MapStringToString(mapFn func(map[string]string) interface{}) refreshable.Refreshable {
	return r.Map(func(i interface{}) interface{} {
		return mapFn(i.(map[string]string))
	})
}

func (r RefreshingStringToString) SubscribeToStringToString(consumer func(map[string]string)) (unsubscribe func()) {
	return r.Subscribe(func(i interface{}) {
		consumer(i.(map[string]string))
	})
}

type RefreshableSecurityConfig interface {
	refreshable.Refreshable
	CurrentSecurityConfig() SecurityConfig
	MapSecurityConfig(func(SecurityConfig) interface{}) refreshable.Refreshable
	SubscribeToSecurityConfig(func(SecurityConfig)) (unsubscribe func())

	CAFiles() refreshable.StringSlice
	CertFile() refreshable.String
	KeyFile() refreshable.String
}

type RefreshingSecurityConfig struct {
	refreshable.Refreshable
}

func NewRefreshingSecurityConfig(in refreshable.Refreshable) RefreshingSecurityConfig {
	return RefreshingSecurityConfig{Refreshable: in}
}

func (r RefreshingSecurityConfig) CurrentSecurityConfig() SecurityConfig {
	return r.Current().(SecurityConfig)
}

func (r RefreshingSecurityConfig) MapSecurityConfig(mapFn func(SecurityConfig) interface{}) refreshable.Refreshable {
	return r.Map(func(i interface{}) interface{} {
		return mapFn(i.(SecurityConfig))
	})
}

func (r RefreshingSecurityConfig) SubscribeToSecurityConfig(consumer func(SecurityConfig)) (unsubscribe func()) {
	return r.Subscribe(func(i interface{}) {
		consumer(i.(SecurityConfig))
	})
}

func (r RefreshingSecurityConfig) CAFiles() refreshable.StringSlice {
	return refreshable.NewStringSlice(r.MapSecurityConfig(func(i SecurityConfig) interface{} {
		return i.CAFiles
	}))
}

func (r RefreshingSecurityConfig) CertFile() refreshable.String {
	return refreshable.NewString(r.MapSecurityConfig(func(i SecurityConfig) interface{} {
		return i.CertFile
	}))
}

func (r RefreshingSecurityConfig) KeyFile() refreshable.String {
	return refreshable.NewString(r.MapSecurityConfig(func(i SecurityConfig) interface{} {
		return i.KeyFile
	}))
}

type RefreshableStringToClientConfig interface {
	refreshable.Refreshable
	CurrentStringToClientConfig() map[string]ClientConfig
	MapStringToClientConfig(func(map[string]ClientConfig) interface{}) refreshable.Refreshable
	SubscribeToStringToClientConfig(func(map[string]ClientConfig)) (unsubscribe func())
}

type RefreshingStringToClientConfig struct {
	refreshable.Refreshable
}

func NewRefreshingStringToClientConfig(in refreshable.Refreshable) RefreshingStringToClientConfig {
	return RefreshingStringToClientConfig{Refreshable: in}
}

func (r RefreshingStringToClientConfig) CurrentStringToClientConfig() map[string]ClientConfig {
	return r.Current().(map[string]ClientConfig)
}

func (r RefreshingStringToClientConfig) MapStringToClientConfig(mapFn func(map[string]ClientConfig) interface{}) refreshable.Refreshable {
	return r.Map(func(i interface{}) interface{} {
		return mapFn(i.(map[string]ClientConfig))
	})
}

func (r RefreshingStringToClientConfig) SubscribeToStringToClientConfig(consumer func(map[string]ClientConfig)) (unsubscribe func()) {
	return r.Subscribe(func(i interface{}) {
		consumer(i.(map[string]ClientConfig))
	})
}
