package entity

type BurpConfig struct {
	Logger struct {
		CaptureFilter struct {
			ByMimeType struct {
				CaptureCss         bool `json:"capture_css"`
				CaptureFlash       bool `json:"capture_flash"`
				CaptureHtml        bool `json:"capture_html"`
				CaptureImages      bool `json:"capture_images"`
				CaptureOtherBinary bool `json:"capture_other_binary"`
				CaptureOtherText   bool `json:"capture_other_text"`
				CaptureScript      bool `json:"capture_script"`
				CaptureXml         bool `json:"capture_xml"`
			} `json:"by_mime_type"`
			ByRequestType struct {
				CaptureOnlyInScopeItems          bool `json:"capture_only_in_scope_items"`
				CaptureOnlyParameterizedRequests bool `json:"capture_only_parameterized_requests"`
				DiscardItemsWithoutResponses     bool `json:"discard_items_without_responses"`
			} `json:"by_request_type"`
			BySearch struct {
				CaseSensitive  bool   `json:"case_sensitive"`
				NegativeSearch bool   `json:"negative_search"`
				Regex          bool   `json:"regex"`
				Term           string `json:"term"`
			} `json:"by_search"`
			ByStatusCode struct {
				Capture2Xx bool `json:"capture_2xx"`
				Capture3Xx bool `json:"capture_3xx"`
				Capture4Xx bool `json:"capture_4xx"`
				Capture5Xx bool `json:"capture_5xx"`
			} `json:"by_status_code"`
			ByTool struct {
				CaptureExtender  bool `json:"capture_extender"`
				CaptureIntruder  bool `json:"capture_intruder"`
				CaptureProxy     bool `json:"capture_proxy"`
				CaptureRepeater  bool `json:"capture_repeater"`
				CaptureScanner   bool `json:"capture_scanner"`
				CaptureSequencer bool `json:"capture_sequencer"`
				CaptureTarget    bool `json:"capture_target"`
			} `json:"by_tool"`
			CaptureEnabled           bool `json:"capture_enabled"`
			CaptureMemoryLimitMb     int  `json:"capture_memory_limit_mb"`
			LimitRequestResponseSize struct {
				CaptureRequestsUpTo  string `json:"capture_requests_up_to"`
				CaptureResponsesUpTo string `json:"capture_responses_up_to"`
			} `json:"limit_request_response_size"`
			SessionHandling struct {
				IgnoreSessionHandlingRequests bool `json:"ignore_session_handling_requests"`
			} `json:"session_handling"`
			TaskCaptureMemoryLimitMb int `json:"task_capture_memory_limit_mb"`
		} `json:"capture_filter"`
		DisplayFilter struct {
			ByAnnotation struct {
				ShowOnlyCommentedItems   bool `json:"show_only_commented_items"`
				ShowOnlyHighlightedItems bool `json:"show_only_highlighted_items"`
			} `json:"by_annotation"`
			ByFileExtension struct {
				HideItems        []string `json:"hide_items"`
				HideSpecific     bool     `json:"hide_specific"`
				ShowItems        []string `json:"show_items"`
				ShowOnlySpecific bool     `json:"show_only_specific"`
			} `json:"by_file_extension"`
			ByMimeType struct {
				ShowCss         bool `json:"show_css"`
				ShowFlash       bool `json:"show_flash"`
				ShowHtml        bool `json:"show_html"`
				ShowImages      bool `json:"show_images"`
				ShowOtherBinary bool `json:"show_other_binary"`
				ShowOtherText   bool `json:"show_other_text"`
				ShowScript      bool `json:"show_script"`
				ShowXml         bool `json:"show_xml"`
			} `json:"by_mime_type"`
			ByRequestType struct {
				HideItemsWithoutResponses     bool `json:"hide_items_without_responses"`
				ShowOnlyInScopeItems          bool `json:"show_only_in_scope_items"`
				ShowOnlyParameterizedRequests bool `json:"show_only_parameterized_requests"`
			} `json:"by_request_type"`
			BySearch struct {
				CaseSensitive  bool   `json:"case_sensitive"`
				NegativeSearch bool   `json:"negative_search"`
				Regex          bool   `json:"regex"`
				Term           string `json:"term"`
			} `json:"by_search"`
			ByStatusCode struct {
				Show2Xx bool `json:"show_2xx"`
				Show3Xx bool `json:"show_3xx"`
				Show4Xx bool `json:"show_4xx"`
				Show5Xx bool `json:"show_5xx"`
			} `json:"by_status_code"`
			ByTool struct {
				ShowExtender  bool `json:"show_extender"`
				ShowIntruder  bool `json:"show_intruder"`
				ShowProxy     bool `json:"show_proxy"`
				ShowRepeater  bool `json:"show_repeater"`
				ShowScanner   bool `json:"show_scanner"`
				ShowSequencer bool `json:"show_sequencer"`
				ShowTarget    bool `json:"show_target"`
			} `json:"by_tool"`
		} `json:"display_filter"`
	} `json:"logger"`
	ProjectOptions struct {
		Connections struct {
			HostnameResolution []interface{} `json:"hostname_resolution"`
			OutOfScopeRequests struct {
				AdvancedMode      bool          `json:"advanced_mode"`
				DropAllOutOfScope bool          `json:"drop_all_out_of_scope"`
				Exclude           []interface{} `json:"exclude"`
				Include           []interface{} `json:"include"`
				ScopeOption       string        `json:"scope_option"`
			} `json:"out_of_scope_requests"`
			PlatformAuthentication struct {
				Credentials                   []interface{} `json:"credentials"`
				DoPlatformAuthentication      bool          `json:"do_platform_authentication"`
				PromptOnAuthenticationFailure bool          `json:"prompt_on_authentication_failure"`
				UseUserOptions                bool          `json:"use_user_options"`
			} `json:"platform_authentication"`
			SocksProxy struct {
				DnsOverSocks   bool   `json:"dns_over_socks"`
				Host           string `json:"host"`
				Password       string `json:"password"`
				Port           int    `json:"port"`
				UseProxy       bool   `json:"use_proxy"`
				UseUserOptions bool   `json:"use_user_options"`
				Username       string `json:"username"`
			} `json:"socks_proxy"`
			Timeouts struct {
				ConnectTimeout                    int `json:"connect_timeout"`
				DomainNameResolutionTimeout       int `json:"domain_name_resolution_timeout"`
				FailedDomainNameResolutionTimeout int `json:"failed_domain_name_resolution_timeout"`
				NormalTimeout                     int `json:"normal_timeout"`
				OpenEndedResponseTimeout          int `json:"open_ended_response_timeout"`
			} `json:"timeouts"`
			UpstreamProxy struct {
				Servers        []interface{} `json:"servers"`
				UseUserOptions bool          `json:"use_user_options"`
			} `json:"upstream_proxy"`
		} `json:"connections"`
		Http struct {
			Http1 struct {
				EnableKeepAlive bool `json:"enable_keep_alive"`
			} `json:"http1"`
			Http2 struct {
				EnableHttp2 bool `json:"enable_http2"`
			} `json:"http2"`
			Redirections struct {
				Understand3XxStatusCode                   bool `json:"understand_3xx_status_code"`
				UnderstandAnyStatusCodeWithLocationHeader bool `json:"understand_any_status_code_with_location_header"`
				UnderstandJavascriptDriven                bool `json:"understand_javascript_driven"`
				UnderstandMetaRefreshTag                  bool `json:"understand_meta_refresh_tag"`
				UnderstandRefreshHeader                   bool `json:"understand_refresh_header"`
			} `json:"redirections"`
			Status100Responses struct {
				Remove100ContinueResponses     bool `json:"remove_100_continue_responses"`
				Understand100ContinueResponses bool `json:"understand_100_continue_responses"`
			} `json:"status_100_responses"`
			StreamingResponses struct {
				ScopeAdvancedMode            bool          `json:"scope_advanced_mode"`
				Store                        bool          `json:"store"`
				StripChunkedEncodingMetadata bool          `json:"strip_chunked_encoding_metadata"`
				Urls                         []interface{} `json:"urls"`
			} `json:"streaming_responses"`
		} `json:"http"`
		Misc struct {
			CollaboratorServer struct {
				Location                string `json:"location"`
				PollOverUnencryptedHttp bool   `json:"poll_over_unencrypted_http"`
				PollingLocation         string `json:"polling_location"`
				Type                    string `json:"type"`
			} `json:"collaborator_server"`
			EmbeddedBrowser struct {
				AllowRunningWithoutSandbox bool `json:"allow_running_without_sandbox"`
				DisableGpu                 bool `json:"disable_gpu"`
			} `json:"embedded_browser"`
			Logging struct {
				Requests struct {
					AllTools  string `json:"all_tools"`
					Extender  string `json:"extender"`
					Intruder  string `json:"intruder"`
					Proxy     string `json:"proxy"`
					Repeater  string `json:"repeater"`
					Scanner   string `json:"scanner"`
					Sequencer string `json:"sequencer"`
				} `json:"requests"`
				Responses struct {
					AllTools  string `json:"all_tools"`
					Extender  string `json:"extender"`
					Intruder  string `json:"intruder"`
					Proxy     string `json:"proxy"`
					Repeater  string `json:"repeater"`
					Scanner   string `json:"scanner"`
					Sequencer string `json:"sequencer"`
				} `json:"responses"`
			} `json:"logging"`
			ScheduledTasks struct {
				Tasks []interface{} `json:"tasks"`
			} `json:"scheduled_tasks"`
		} `json:"misc"`
		Sessions struct {
			CookieJar struct {
				MonitorExtender  bool `json:"monitor_extender"`
				MonitorIntruder  bool `json:"monitor_intruder"`
				MonitorProxy     bool `json:"monitor_proxy"`
				MonitorRepeater  bool `json:"monitor_repeater"`
				MonitorScanner   bool `json:"monitor_scanner"`
				MonitorSequencer bool `json:"monitor_sequencer"`
			} `json:"cookie_jar"`
			Macros struct {
				Macros []interface{} `json:"macros"`
			} `json:"macros"`
			SessionHandlingRules struct {
				Rules []struct {
					Actions []struct {
						Enabled      bool   `json:"enabled"`
						MatchCookies string `json:"match_cookies"`
						Type         string `json:"type"`
					} `json:"actions"`
					Description                string        `json:"description"`
					Enabled                    bool          `json:"enabled"`
					ExcludeFromScope           []interface{} `json:"exclude_from_scope"`
					IncludeInScope             []interface{} `json:"include_in_scope"`
					NamedParams                []interface{} `json:"named_params"`
					RestrictScopeToNamedParams bool          `json:"restrict_scope_to_named_params"`
					ToolsScope                 []string      `json:"tools_scope"`
					UrlScope                   string        `json:"url_scope"`
					UrlScopeAdvancedMode       bool          `json:"url_scope_advanced_mode"`
				} `json:"rules"`
			} `json:"session_handling_rules"`
		} `json:"sessions"`
		Ssl struct {
			ClientCertificates struct {
				Certificates   []interface{} `json:"certificates"`
				UseUserOptions bool          `json:"use_user_options"`
			} `json:"client_certificates"`
			Negotiation struct {
				AllowUnsafeRenegotiation bool          `json:"allow_unsafe_renegotiation"`
				DisableSslSessionResume  bool          `json:"disable_ssl_session_resume"`
				EnabledCiphers           []interface{} `json:"enabled_ciphers"`
				EnabledProtocols         []interface{} `json:"enabled_protocols"`
				EnforceUpstreamTrust     bool          `json:"enforce_upstream_trust"`
				TlsNegotiationBehavior   string        `json:"tls_negotiation_behavior"`
			} `json:"negotiation"`
		} `json:"ssl"`
	} `json:"project_options"`
	Proxy struct {
		HttpHistoryDisplayFilter struct {
			ByAnnotation struct {
				ShowOnlyCommentedItems   bool `json:"show_only_commented_items"`
				ShowOnlyHighlightedItems bool `json:"show_only_highlighted_items"`
			} `json:"by_annotation"`
			ByFileExtension struct {
				HideItems        []string `json:"hide_items"`
				HideSpecific     bool     `json:"hide_specific"`
				ShowItems        []string `json:"show_items"`
				ShowOnlySpecific bool     `json:"show_only_specific"`
			} `json:"by_file_extension"`
			ByListener struct {
				Port string `json:"port"`
			} `json:"by_listener"`
			ByMimeType struct {
				ShowCss         bool `json:"show_css"`
				ShowFlash       bool `json:"show_flash"`
				ShowHtml        bool `json:"show_html"`
				ShowImages      bool `json:"show_images"`
				ShowOtherBinary bool `json:"show_other_binary"`
				ShowOtherText   bool `json:"show_other_text"`
				ShowScript      bool `json:"show_script"`
				ShowXml         bool `json:"show_xml"`
			} `json:"by_mime_type"`
			ByRequestType struct {
				HideItemsWithoutResponses     bool `json:"hide_items_without_responses"`
				ShowOnlyInScopeItems          bool `json:"show_only_in_scope_items"`
				ShowOnlyParameterizedRequests bool `json:"show_only_parameterized_requests"`
			} `json:"by_request_type"`
			BySearch struct {
				CaseSensitive  bool   `json:"case_sensitive"`
				NegativeSearch bool   `json:"negative_search"`
				Regex          bool   `json:"regex"`
				Term           string `json:"term"`
			} `json:"by_search"`
			ByStatusCode struct {
				Show2Xx bool `json:"show_2xx"`
				Show3Xx bool `json:"show_3xx"`
				Show4Xx bool `json:"show_4xx"`
				Show5Xx bool `json:"show_5xx"`
			} `json:"by_status_code"`
		} `json:"http_history_display_filter"`
		InterceptClientRequests struct {
			AutomaticallyFixMissingOrSuperfluousNewLinesAtEndOfRequest   bool `json:"automatically_fix_missing_or_superfluous_new_lines_at_end_of_request"`
			AutomaticallyUpdateContentLengthHeaderWhenTheRequestIsEdited bool `json:"automatically_update_content_length_header_when_the_request_is_edited"`
			DoIntercept                                                  bool `json:"do_intercept"`
			Rules                                                        []struct {
				BooleanOperator   string `json:"boolean_operator"`
				Enabled           bool   `json:"enabled"`
				MatchCondition    string `json:"match_condition,omitempty"`
				MatchRelationship string `json:"match_relationship"`
				MatchType         string `json:"match_type"`
			} `json:"rules"`
		} `json:"intercept_client_requests"`
		InterceptServerResponses struct {
			AutomaticallyUpdateContentLengthHeaderWhenTheResponseIsEdited bool `json:"automatically_update_content_length_header_when_the_response_is_edited"`
			DoIntercept                                                   bool `json:"do_intercept"`
			Rules                                                         []struct {
				BooleanOperator   string `json:"boolean_operator"`
				Enabled           bool   `json:"enabled"`
				MatchCondition    string `json:"match_condition,omitempty"`
				MatchRelationship string `json:"match_relationship"`
				MatchType         string `json:"match_type"`
			} `json:"rules"`
		} `json:"intercept_server_responses"`
		InterceptWebSocketsMessages struct {
			ClientToServerMessages bool `json:"client_to_server_messages"`
			ServerToClientMessages bool `json:"server_to_client_messages"`
		} `json:"intercept_web_sockets_messages"`
		MatchReplaceRules []struct {
			Comment       string `json:"comment"`
			Enabled       bool   `json:"enabled"`
			IsSimpleMatch bool   `json:"is_simple_match"`
			RuleType      string `json:"rule_type"`
			StringMatch   string `json:"string_match,omitempty"`
			StringReplace string `json:"string_replace,omitempty"`
		} `json:"match_replace_rules"`
		Miscellaneous struct {
			DisableLoggingToHistoryAndSiteMap                                     bool `json:"disable_logging_to_history_and_site_map"`
			DisableOutOfScopeLoggingToHistoryAndSiteMap                           bool `json:"disable_out_of_scope_logging_to_history_and_site_map"`
			DisableWebInterface                                                   bool `json:"disable_web_interface"`
			RemoveUnsupportedEncodingsFromAcceptEncodingHeadersInIncomingRequests bool `json:"remove_unsupported_encodings_from_accept_encoding_headers_in_incoming_requests"`
			SetConnectionCloseHeaderOnResponses                                   bool `json:"set_connection_close_header_on_responses"`
			SetConnectionHeaderOnRequests                                         bool `json:"set_connection_header_on_requests"`
			StripProxyHeadersInIncomingRequests                                   bool `json:"strip_proxy_headers_in_incoming_requests"`
			StripSecWebsocketExtensionsHeadersInIncomingRequests                  bool `json:"strip_sec_websocket_extensions_headers_in_incoming_requests"`
			SuppressBurpErrorMessagesInBrowser                                    bool `json:"suppress_burp_error_messages_in_browser"`
			UnpackGzipDeflateInRequests                                           bool `json:"unpack_gzip_deflate_in_requests"`
			UnpackGzipDeflateInResponses                                          bool `json:"unpack_gzip_deflate_in_responses"`
			UseHttp10InRequestsToServer                                           bool `json:"use_http_10_in_requests_to_server"`
			UseHttp10InResponsesToClient                                          bool `json:"use_http_10_in_responses_to_client"`
		} `json:"miscellaneous"`
		RequestListeners []struct {
			CertificateMode       string        `json:"certificate_mode"`
			CustomTlsProtocols    []interface{} `json:"custom_tls_protocols"`
			EnableHttp2           bool          `json:"enable_http2"`
			ListenMode            string        `json:"listen_mode"`
			ListenerPort          int           `json:"listener_port"`
			Running               bool          `json:"running"`
			UseCustomTlsProtocols bool          `json:"use_custom_tls_protocols"`
		} `json:"request_listeners"`
		ResponseModification struct {
			ConvertHttpsLinksToHttp        bool `json:"convert_https_links_to_http"`
			EnableDisabledFormFields       bool `json:"enable_disabled_form_fields"`
			HighlightUnhiddenFields        bool `json:"highlight_unhidden_fields"`
			RemoveAllJavascript            bool `json:"remove_all_javascript"`
			RemoveInputFieldLengthLimits   bool `json:"remove_input_field_length_limits"`
			RemoveJavascriptFormValidation bool `json:"remove_javascript_form_validation"`
			RemoveObjectTags               bool `json:"remove_object_tags"`
			RemoveSecureFlagFromCookies    bool `json:"remove_secure_flag_from_cookies"`
			UnhideHiddenFormFields         bool `json:"unhide_hidden_form_fields"`
		} `json:"response_modification"`
		SslPassThrough struct {
			AutomaticallyAddEntriesOnClientSslNegotiationFailure bool          `json:"automatically_add_entries_on_client_ssl_negotiation_failure"`
			Rules                                                []interface{} `json:"rules"`
		} `json:"ssl_pass_through"`
		WebSocketsHistoryDisplayFilter struct {
			ByAnnotation struct {
				ShowOnlyCommentedItems   bool `json:"show_only_commented_items"`
				ShowOnlyHighlightedItems bool `json:"show_only_highlighted_items"`
			} `json:"by_annotation"`
			ByListener struct {
				ListenerPort string `json:"listener_port"`
			} `json:"by_listener"`
			ByRequestType struct {
				HideIncomingMessages bool `json:"hide_incoming_messages"`
				HideOutgoingMessages bool `json:"hide_outgoing_messages"`
				ShowOnlyInScopeItems bool `json:"show_only_in_scope_items"`
			} `json:"by_request_type"`
			BySearch struct {
				CaseSensitive  bool   `json:"case_sensitive"`
				NegativeSearch bool   `json:"negative_search"`
				Regex          bool   `json:"regex"`
				Term           string `json:"term"`
			} `json:"by_search"`
		} `json:"web_sockets_history_display_filter"`
	} `json:"proxy"`
	Repeater struct {
		AllowHttp2AlpnOverride         bool   `json:"allow_http2_alpn_override"`
		EnableHttp1KeepAlive           bool   `json:"enable_http1_keep_alive"`
		EnableHttp2ConnectionReuse     bool   `json:"enable_http2_connection_reuse"`
		EnforceProtocolInRedirections  bool   `json:"enforce_protocol_in_redirections"`
		FollowRedirections             string `json:"follow_redirections"`
		NormalizeLineEndings           bool   `json:"normalize_line_endings"`
		ProcessCookiesInRedirections   bool   `json:"process_cookies_in_redirections"`
		StripConnectionHeaderOverHttp2 bool   `json:"strip_connection_header_over_http2"`
		UnpackGzipDeflate              bool   `json:"unpack_gzip_deflate"`
		UpdateContentLength            bool   `json:"update_content_length"`
	} `json:"repeater"`
	Sequencer struct {
		LiveCapture struct {
			IgnoreAbnormalLengthTokens bool `json:"ignore_abnormal_length_tokens"`
			MaxLengthDeviation         int  `json:"max_length_deviation"`
			NumThreads                 int  `json:"num_threads"`
			Throttle                   int  `json:"throttle"`
		} `json:"live_capture"`
		TokenAnalysis struct {
			Compression bool `json:"compression"`
			Correlation bool `json:"correlation"`
			Count       bool `json:"count"`
			FipsLongRun bool `json:"fips_long_run"`
			FipsMonobit bool `json:"fips_monobit"`
			FipsPoker   bool `json:"fips_poker"`
			FipsRuns    bool `json:"fips_runs"`
			Spectral    bool `json:"spectral"`
			Transitions bool `json:"transitions"`
		} `json:"token_analysis"`
		TokenHandling struct {
			Base64DecodeBeforeAnalyzing bool   `json:"base_64_decode_before_analyzing"`
			PadShortTokensAt            string `json:"pad_short_tokens_at"`
			PadWith                     string `json:"pad_with"`
		} `json:"token_handling"`
	} `json:"sequencer"`
	Target struct {
		Filter struct {
			ByAnnotation struct {
				ShowOnlyCommentedItems   bool `json:"show_only_commented_items"`
				ShowOnlyHighlightedItems bool `json:"show_only_highlighted_items"`
			} `json:"by_annotation"`
			ByFileExtension struct {
				HideItems        []string `json:"hide_items"`
				HideSpecific     bool     `json:"hide_specific"`
				ShowItems        []string `json:"show_items"`
				ShowOnlySpecific bool     `json:"show_only_specific"`
			} `json:"by_file_extension"`
			ByFolders struct {
				HideEmptyFolders bool `json:"hide_empty_folders"`
			} `json:"by_folders"`
			ByMimeType struct {
				ShowCss         bool `json:"show_css"`
				ShowFlash       bool `json:"show_flash"`
				ShowHtml        bool `json:"show_html"`
				ShowImages      bool `json:"show_images"`
				ShowOtherBinary bool `json:"show_other_binary"`
				ShowOtherText   bool `json:"show_other_text"`
				ShowScript      bool `json:"show_script"`
				ShowXml         bool `json:"show_xml"`
			} `json:"by_mime_type"`
			ByRequestType struct {
				HideNotFoundItems             bool `json:"hide_not_found_items"`
				ShowOnlyInScopeItems          bool `json:"show_only_in_scope_items"`
				ShowOnlyParameterizedRequests bool `json:"show_only_parameterized_requests"`
				ShowOnlyRequestedItems        bool `json:"show_only_requested_items"`
			} `json:"by_request_type"`
			BySearch struct {
				CaseSensitive  bool   `json:"case_sensitive"`
				NegativeSearch bool   `json:"negative_search"`
				Regex          bool   `json:"regex"`
				Term           string `json:"term"`
			} `json:"by_search"`
			ByStatusCode struct {
				Show2Xx bool `json:"show_2xx"`
				Show3Xx bool `json:"show_3xx"`
				Show4Xx bool `json:"show_4xx"`
				Show5Xx bool `json:"show_5xx"`
			} `json:"by_status_code"`
		} `json:"filter"`
		Scope struct {
			AdvancedMode bool          `json:"advanced_mode"`
			Exclude      []interface{} `json:"exclude"`
			Include      []interface{} `json:"include"`
		} `json:"scope"`
	} `json:"target"`
}
