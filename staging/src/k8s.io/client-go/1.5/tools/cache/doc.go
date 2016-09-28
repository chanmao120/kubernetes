/*
Copyright 2014 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Package cache is a client-side caching mechanism. It is useful for
// reducing the number of server calls you'd otherwise need to make.
// Reflector watches a server and updates a Store. Two stores are provided;
// one that simply caches objects (for example, to allow a scheduler to
// list currently available nodes), and one that additionally acts as
// a FIFO queue (for example, to allow a scheduler to process incoming
// pods).
package testing // import "k8s.io/client-go/1.5/tools/cache"