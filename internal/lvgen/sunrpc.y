%{
// Copyright 2017 The go-libvirt Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package lvgen

import (
    //"fmt"
)

%}

// SymType
%union{
    val string
}

// XDR tokens:
%token BOOL CASE CONST DEFAULT DOUBLE ENUM FLOAT OPAQUE STRING STRUCT
%token SWITCH TYPEDEF UNION UNSIGNED VOID HYPER INT SHORT CHAR
%token IDENTIFIER CONSTANT ERROR
// RPCL additional tokens:
%token PROGRAM VERSION

%%

specification
    : definition_list
    ;

value
    : IDENTIFIER
    | CONSTANT
    ;

definition_list
    : definition ';'
    | definition ';' definition_list
    ;

definition
    : enum_definition
    | const_definition
    | typedef_definition
    | struct_definition
    | union_definition
    | program_definition
    ;

enum_definition
    : ENUM enum_ident '{' enum_value_list '}' { StartEnum($2.val) }
    ;

enum_value_list
    : enum_value
    | enum_value ',' enum_value_list
    ;

enum_value
    : enum_value_ident {
        err := AddEnumAutoVal($1.val)
        if err != nil {
            yylex.Error(err.Error())
            return 1
        }
    }
    | enum_value_ident '=' value {
        err := AddEnumVal($1.val, $3.val)
        if err != nil {
            yylex.Error(err.Error())
            return 1
        }
    }
    ;

enum_ident
    : IDENTIFIER
    ;

enum_value_ident
    : IDENTIFIER
    ;

// Ignore consts that are set to IDENTIFIERs - this isn't allowed by the spec,
// but occurs in the file because libvirt runs the pre-processor on the protocol
// file, and it handles replacing the identifier with it's #defined value.
const_definition
    : CONST const_ident '=' IDENTIFIER
    | CONST const_ident '=' CONSTANT {
        err := AddConst($2.val, $4.val)
        if err != nil {
            yylex.Error(err.Error())
            return 1
        }
    }
    ;

const_ident
    : IDENTIFIER
    ;

typedef_definition
    : TYPEDEF {StartTypedef()} declaration
    ;

declaration
    : simple_declaration
    | fixed_array_declaration
    | variable_array_declaration
    | pointer_declaration
    ;

simple_declaration
    : type_specifier variable_ident {AddDeclaration($2.val, $1.val)}
    ;

type_specifier
    : int_spec
    | UNSIGNED int_spec {$$.val = "u"+$2.val}
    | FLOAT             {$$.val = "float32"}
    | DOUBLE            {$$.val = "float64"}
    | BOOL              {$$.val = "bool"}
    | STRING            {$$.val = "string"}
    | OPAQUE            {$$.val = "byte"}
    | enum_definition
    | struct_definition
    | union_definition
    | IDENTIFIER
    ;

int_spec
    : HYPER {$$.val = "int64"}
    | INT   {$$.val = "int32"}
    | SHORT {$$.val = "int16"}
    | CHAR  {$$.val = "int8"}
    ;

variable_ident
    : IDENTIFIER
    ;

fixed_array_declaration
    : type_specifier variable_ident '[' value ']'   { AddFixedArray($2.val, $1.val, $4.val) }
    ;

variable_array_declaration
    : type_specifier variable_ident '<' value '>'   { AddVariableArray($2.val, $1.val, $4.val) }
    | type_specifier variable_ident '<' '>'         { AddVariableArray($2.val, $1.val, "") }
    ;

// while pointer_declarations may look like their familiar c-equivalents, in the
// XDR language they actually declare "Optional-data". The simplest
// representation to use for these is a variable-length array with a size of 1.
// See the XDR spec for a more complete explanation of this.
pointer_declaration
    : type_specifier '*' variable_ident             { AddVariableArray($3.val, $1.val, "1") }
    ;

struct_definition
    : STRUCT struct_ident '{' {StartStruct($2.val)} declaration_list '}' { AddStruct() }
    ;

struct_ident
    : IDENTIFIER
    ;

declaration_list
    : declaration ';'
    | declaration ';' declaration_list
    ;

union_definition
    : UNION union_ident {StartUnion($2.val)} SWITCH '(' simple_declaration ')' '{' case_list '}' {AddUnion()}
    ;

union_ident
    : IDENTIFIER
    ;

case_list
    : case ';'
    | case ';' case_list
    ;

case
    : CASE value {StartCase($2.val)} ':' declaration {AddCase()}
    | DEFAULT {StartCase("default")} ':' declaration {AddCase()}
    ;

program_definition
    : PROGRAM program_ident '{' version_list '}' '=' value
    ;

program_ident
    : IDENTIFIER
    ;

version_list
    : version ';'
    | version ';' version_list
    ;

version
    : VERSION version_ident '{' procedure_list '}' '=' value ';'
    ;

version_ident
    : IDENTIFIER
    ;

procedure_list
    : procedure ';'
    | procedure ';' procedure_list
    ;

procedure
    : type_specifier procedure_ident '(' type_specifier ')' '=' value ';'
    ;

procedure_ident
    : IDENTIFIER
    ;

%%
