#
# Copyright (c) 2020-present, Weibo, Inc.  All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# @author ZhongXiu Hao <nmred.hao@gmail.com>

load("@//thirdparty/tools:fbthrift.bzl", "gen_fbthrift")

gen_fbthrift(
  name = "geo",
  src = "geo/if/geo.thrift",
  output_path = "geo/if",
  service_list = [
    "GeoService",
  ],
  options = "templates,no_metadata",
)

cc_binary(
	name = "geo_server",
	srcs=[
		"geo/server/geo_service.h",
		"geo/server/geo_service.cc",
		"geo/server/geo_server.cc",
    ":geo",
	],
	deps=[
    "//common:util",
    '//common:service_router_http_server',
		'@fbthrift//:thriftcpp2',
	],
  includes = [
    "."
  ],
  linkopts = [
    '-ldl'
  ],
  copts = [
    "-Iexternal/double-conversion/",
  ],
)

cc_binary(
  name = "geo_client",
  srcs = [
    "geo/client/geo_client.cc",
    ":geo",
  ],
  deps = [
    "@folly//:folly",
    "//common:service_router_http_server",
		"@fbthrift//:thriftcpp2",
  ],
  includes = [
    "."
  ],
  linkopts = [
    '-ldl'
  ],
  copts = [
    "-Iexternal/double-conversion/",
  ],
)
