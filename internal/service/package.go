// See the file LICENSE for redistribution and license information.
//
// Copyright (c) 2020 worldiety. All rights reserved.
// DO NOT ALTER OR REMOVE COPYRIGHT NOTICES OR THIS FILE HEADER.
//
// Please contact worldiety, Marie-Curie-Straße 1, 26129 Oldenburg, Germany
// or visit www.worldiety.com if you need more information or have any questions.
//
// Authors: Torben Schinke

// Package service contains all internal services, splitted by a variant of domain driven design
// which each itself follows clean architecture recommendations.
//
// Services must not have direct dependencies between each other. However if required, they must be
// decoupled by using custom interfaces. Keep in mind, that each service may be pulled into
// its own microservice implementation.
package service
