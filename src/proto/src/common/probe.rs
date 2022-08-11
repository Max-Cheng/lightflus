// This file is generated by rust-protobuf 2.27.1. Do not edit
// @generated

// https://github.com/rust-lang/rust-clippy/issues/702
#![allow(unknown_lints)]
#![allow(clippy::all)]

#![allow(unused_attributes)]
#![cfg_attr(rustfmt, rustfmt::skip)]

#![allow(box_pointers)]
#![allow(dead_code)]
#![allow(missing_docs)]
#![allow(non_camel_case_types)]
#![allow(non_snake_case)]
#![allow(non_upper_case_globals)]
#![allow(trivial_casts)]
#![allow(unused_imports)]
#![allow(unused_results)]
//! Generated file from `common/probe.proto`

/// Generated files are compatible only with the same version
/// of protobuf runtime.
// const _PROTOBUF_VERSION_CHECK: () = ::protobuf::VERSION_2_27_1;

#[derive(PartialEq,Clone,Default)]
pub struct ProbeRequest {
    // message fields
    pub nodeType: ProbeRequest_NodeType,
    pub probeType: ProbeRequest_ProbeType,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a ProbeRequest {
    fn default() -> &'a ProbeRequest {
        <ProbeRequest as ::protobuf::Message>::default_instance()
    }
}

impl ProbeRequest {
    pub fn new() -> ProbeRequest {
        ::std::default::Default::default()
    }

    // .common.ProbeRequest.NodeType nodeType = 1;


    pub fn get_nodeType(&self) -> ProbeRequest_NodeType {
        self.nodeType
    }
    pub fn clear_nodeType(&mut self) {
        self.nodeType = ProbeRequest_NodeType::Coordinator;
    }

    // Param is passed by value, moved
    pub fn set_nodeType(&mut self, v: ProbeRequest_NodeType) {
        self.nodeType = v;
    }

    // .common.ProbeRequest.ProbeType probeType = 2;


    pub fn get_probeType(&self) -> ProbeRequest_ProbeType {
        self.probeType
    }
    pub fn clear_probeType(&mut self) {
        self.probeType = ProbeRequest_ProbeType::Liveness;
    }

    // Param is passed by value, moved
    pub fn set_probeType(&mut self, v: ProbeRequest_ProbeType) {
        self.probeType = v;
    }
}

impl ::protobuf::Message for ProbeRequest {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_proto3_enum_with_unknown_fields_into(wire_type, is, &mut self.nodeType, 1, &mut self.unknown_fields)?
                },
                2 => {
                    ::protobuf::rt::read_proto3_enum_with_unknown_fields_into(wire_type, is, &mut self.probeType, 2, &mut self.unknown_fields)?
                },
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        if self.nodeType != ProbeRequest_NodeType::Coordinator {
            my_size += ::protobuf::rt::enum_size(1, self.nodeType);
        }
        if self.probeType != ProbeRequest_ProbeType::Liveness {
            my_size += ::protobuf::rt::enum_size(2, self.probeType);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if self.nodeType != ProbeRequest_NodeType::Coordinator {
            os.write_enum(1, ::protobuf::ProtobufEnum::value(&self.nodeType))?;
        }
        if self.probeType != ProbeRequest_ProbeType::Liveness {
            os.write_enum(2, ::protobuf::ProtobufEnum::value(&self.probeType))?;
        }
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> ProbeRequest {
        ProbeRequest::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeEnum<ProbeRequest_NodeType>>(
                "nodeType",
                |m: &ProbeRequest| { &m.nodeType },
                |m: &mut ProbeRequest| { &mut m.nodeType },
            ));
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeEnum<ProbeRequest_ProbeType>>(
                "probeType",
                |m: &ProbeRequest| { &m.probeType },
                |m: &mut ProbeRequest| { &mut m.probeType },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<ProbeRequest>(
                "ProbeRequest",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static ProbeRequest {
        static instance: ::protobuf::rt::LazyV2<ProbeRequest> = ::protobuf::rt::LazyV2::INIT;
        instance.get(ProbeRequest::new)
    }
}

impl ::protobuf::Clear for ProbeRequest {
    fn clear(&mut self) {
        self.nodeType = ProbeRequest_NodeType::Coordinator;
        self.probeType = ProbeRequest_ProbeType::Liveness;
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for ProbeRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for ProbeRequest {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(Clone,PartialEq,Eq,Debug,Hash)]
pub enum ProbeRequest_NodeType {
    Coordinator = 0,
    TaskWorker = 1,
    Connector = 2,
}

impl ::protobuf::ProtobufEnum for ProbeRequest_NodeType {
    fn value(&self) -> i32 {
        *self as i32
    }

    fn from_i32(value: i32) -> ::std::option::Option<ProbeRequest_NodeType> {
        match value {
            0 => ::std::option::Option::Some(ProbeRequest_NodeType::Coordinator),
            1 => ::std::option::Option::Some(ProbeRequest_NodeType::TaskWorker),
            2 => ::std::option::Option::Some(ProbeRequest_NodeType::Connector),
            _ => ::std::option::Option::None
        }
    }

    fn values() -> &'static [Self] {
        static values: &'static [ProbeRequest_NodeType] = &[
            ProbeRequest_NodeType::Coordinator,
            ProbeRequest_NodeType::TaskWorker,
            ProbeRequest_NodeType::Connector,
        ];
        values
    }

    fn enum_descriptor_static() -> &'static ::protobuf::reflect::EnumDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::EnumDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            ::protobuf::reflect::EnumDescriptor::new_pb_name::<ProbeRequest_NodeType>("ProbeRequest.NodeType", file_descriptor_proto())
        })
    }
}

impl ::std::marker::Copy for ProbeRequest_NodeType {
}

impl ::std::default::Default for ProbeRequest_NodeType {
    fn default() -> Self {
        ProbeRequest_NodeType::Coordinator
    }
}

impl ::protobuf::reflect::ProtobufValue for ProbeRequest_NodeType {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Enum(::protobuf::ProtobufEnum::descriptor(self))
    }
}

#[derive(Clone,PartialEq,Eq,Debug,Hash)]
pub enum ProbeRequest_ProbeType {
    Liveness = 0,
    Readiness = 1,
}

impl ::protobuf::ProtobufEnum for ProbeRequest_ProbeType {
    fn value(&self) -> i32 {
        *self as i32
    }

    fn from_i32(value: i32) -> ::std::option::Option<ProbeRequest_ProbeType> {
        match value {
            0 => ::std::option::Option::Some(ProbeRequest_ProbeType::Liveness),
            1 => ::std::option::Option::Some(ProbeRequest_ProbeType::Readiness),
            _ => ::std::option::Option::None
        }
    }

    fn values() -> &'static [Self] {
        static values: &'static [ProbeRequest_ProbeType] = &[
            ProbeRequest_ProbeType::Liveness,
            ProbeRequest_ProbeType::Readiness,
        ];
        values
    }

    fn enum_descriptor_static() -> &'static ::protobuf::reflect::EnumDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::EnumDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            ::protobuf::reflect::EnumDescriptor::new_pb_name::<ProbeRequest_ProbeType>("ProbeRequest.ProbeType", file_descriptor_proto())
        })
    }
}

impl ::std::marker::Copy for ProbeRequest_ProbeType {
}

impl ::std::default::Default for ProbeRequest_ProbeType {
    fn default() -> Self {
        ProbeRequest_ProbeType::Liveness
    }
}

impl ::protobuf::reflect::ProtobufValue for ProbeRequest_ProbeType {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Enum(::protobuf::ProtobufEnum::descriptor(self))
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct ProbeResponse {
    // message fields
    pub memory: f32,
    pub cpu: f32,
    pub available: bool,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a ProbeResponse {
    fn default() -> &'a ProbeResponse {
        <ProbeResponse as ::protobuf::Message>::default_instance()
    }
}

impl ProbeResponse {
    pub fn new() -> ProbeResponse {
        ::std::default::Default::default()
    }

    // float memory = 1;


    pub fn get_memory(&self) -> f32 {
        self.memory
    }
    pub fn clear_memory(&mut self) {
        self.memory = 0.;
    }

    // Param is passed by value, moved
    pub fn set_memory(&mut self, v: f32) {
        self.memory = v;
    }

    // float cpu = 2;


    pub fn get_cpu(&self) -> f32 {
        self.cpu
    }
    pub fn clear_cpu(&mut self) {
        self.cpu = 0.;
    }

    // Param is passed by value, moved
    pub fn set_cpu(&mut self, v: f32) {
        self.cpu = v;
    }

    // bool available = 3;


    pub fn get_available(&self) -> bool {
        self.available
    }
    pub fn clear_available(&mut self) {
        self.available = false;
    }

    // Param is passed by value, moved
    pub fn set_available(&mut self, v: bool) {
        self.available = v;
    }
}

impl ::protobuf::Message for ProbeResponse {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    if wire_type != ::protobuf::wire_format::WireTypeFixed32 {
                        return ::std::result::Result::Err(::protobuf::rt::unexpected_wire_type(wire_type));
                    }
                    let tmp = is.read_float()?;
                    self.memory = tmp;
                },
                2 => {
                    if wire_type != ::protobuf::wire_format::WireTypeFixed32 {
                        return ::std::result::Result::Err(::protobuf::rt::unexpected_wire_type(wire_type));
                    }
                    let tmp = is.read_float()?;
                    self.cpu = tmp;
                },
                3 => {
                    if wire_type != ::protobuf::wire_format::WireTypeVarint {
                        return ::std::result::Result::Err(::protobuf::rt::unexpected_wire_type(wire_type));
                    }
                    let tmp = is.read_bool()?;
                    self.available = tmp;
                },
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        if self.memory != 0. {
            my_size += 5;
        }
        if self.cpu != 0. {
            my_size += 5;
        }
        if self.available != false {
            my_size += 2;
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if self.memory != 0. {
            os.write_float(1, self.memory)?;
        }
        if self.cpu != 0. {
            os.write_float(2, self.cpu)?;
        }
        if self.available != false {
            os.write_bool(3, self.available)?;
        }
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> ProbeResponse {
        ProbeResponse::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeFloat>(
                "memory",
                |m: &ProbeResponse| { &m.memory },
                |m: &mut ProbeResponse| { &mut m.memory },
            ));
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeFloat>(
                "cpu",
                |m: &ProbeResponse| { &m.cpu },
                |m: &mut ProbeResponse| { &mut m.cpu },
            ));
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeBool>(
                "available",
                |m: &ProbeResponse| { &m.available },
                |m: &mut ProbeResponse| { &mut m.available },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<ProbeResponse>(
                "ProbeResponse",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static ProbeResponse {
        static instance: ::protobuf::rt::LazyV2<ProbeResponse> = ::protobuf::rt::LazyV2::INIT;
        instance.get(ProbeResponse::new)
    }
}

impl ::protobuf::Clear for ProbeResponse {
    fn clear(&mut self) {
        self.memory = 0.;
        self.cpu = 0.;
        self.available = false;
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for ProbeResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for ProbeResponse {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct EventRequest {
    // message fields
    pub data: ::std::vec::Vec<u8>,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a EventRequest {
    fn default() -> &'a EventRequest {
        <EventRequest as ::protobuf::Message>::default_instance()
    }
}

impl EventRequest {
    pub fn new() -> EventRequest {
        ::std::default::Default::default()
    }

    // bytes data = 1;


    pub fn get_data(&self) -> &[u8] {
        &self.data
    }
    pub fn clear_data(&mut self) {
        self.data.clear();
    }

    // Param is passed by value, moved
    pub fn set_data(&mut self, v: ::std::vec::Vec<u8>) {
        self.data = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_data(&mut self) -> &mut ::std::vec::Vec<u8> {
        &mut self.data
    }

    // Take field
    pub fn take_data(&mut self) -> ::std::vec::Vec<u8> {
        ::std::mem::replace(&mut self.data, ::std::vec::Vec::new())
    }
}

impl ::protobuf::Message for EventRequest {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    ::protobuf::rt::read_singular_proto3_bytes_into(wire_type, is, &mut self.data)?;
                },
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        if !self.data.is_empty() {
            my_size += ::protobuf::rt::bytes_size(1, &self.data);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if !self.data.is_empty() {
            os.write_bytes(1, &self.data)?;
        }
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> EventRequest {
        EventRequest::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeBytes>(
                "data",
                |m: &EventRequest| { &m.data },
                |m: &mut EventRequest| { &mut m.data },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<EventRequest>(
                "EventRequest",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static EventRequest {
        static instance: ::protobuf::rt::LazyV2<EventRequest> = ::protobuf::rt::LazyV2::INIT;
        instance.get(EventRequest::new)
    }
}

impl ::protobuf::Clear for EventRequest {
    fn clear(&mut self) {
        self.data.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for EventRequest {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for EventRequest {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

#[derive(PartialEq,Clone,Default)]
pub struct EventResponse {
    // message fields
    pub code: i32,
    pub msg: ::std::string::String,
    // special fields
    pub unknown_fields: ::protobuf::UnknownFields,
    pub cached_size: ::protobuf::CachedSize,
}

impl<'a> ::std::default::Default for &'a EventResponse {
    fn default() -> &'a EventResponse {
        <EventResponse as ::protobuf::Message>::default_instance()
    }
}

impl EventResponse {
    pub fn new() -> EventResponse {
        ::std::default::Default::default()
    }

    // int32 code = 1;


    pub fn get_code(&self) -> i32 {
        self.code
    }
    pub fn clear_code(&mut self) {
        self.code = 0;
    }

    // Param is passed by value, moved
    pub fn set_code(&mut self, v: i32) {
        self.code = v;
    }

    // string msg = 2;


    pub fn get_msg(&self) -> &str {
        &self.msg
    }
    pub fn clear_msg(&mut self) {
        self.msg.clear();
    }

    // Param is passed by value, moved
    pub fn set_msg(&mut self, v: ::std::string::String) {
        self.msg = v;
    }

    // Mutable pointer to the field.
    // If field is not initialized, it is initialized with default value first.
    pub fn mut_msg(&mut self) -> &mut ::std::string::String {
        &mut self.msg
    }

    // Take field
    pub fn take_msg(&mut self) -> ::std::string::String {
        ::std::mem::replace(&mut self.msg, ::std::string::String::new())
    }
}

impl ::protobuf::Message for EventResponse {
    fn is_initialized(&self) -> bool {
        true
    }

    fn merge_from(&mut self, is: &mut ::protobuf::CodedInputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        while !is.eof()? {
            let (field_number, wire_type) = is.read_tag_unpack()?;
            match field_number {
                1 => {
                    if wire_type != ::protobuf::wire_format::WireTypeVarint {
                        return ::std::result::Result::Err(::protobuf::rt::unexpected_wire_type(wire_type));
                    }
                    let tmp = is.read_int32()?;
                    self.code = tmp;
                },
                2 => {
                    ::protobuf::rt::read_singular_proto3_string_into(wire_type, is, &mut self.msg)?;
                },
                _ => {
                    ::protobuf::rt::read_unknown_or_skip_group(field_number, wire_type, is, self.mut_unknown_fields())?;
                },
            };
        }
        ::std::result::Result::Ok(())
    }

    // Compute sizes of nested messages
    #[allow(unused_variables)]
    fn compute_size(&self) -> u32 {
        let mut my_size = 0;
        if self.code != 0 {
            my_size += ::protobuf::rt::value_size(1, self.code, ::protobuf::wire_format::WireTypeVarint);
        }
        if !self.msg.is_empty() {
            my_size += ::protobuf::rt::string_size(2, &self.msg);
        }
        my_size += ::protobuf::rt::unknown_fields_size(self.get_unknown_fields());
        self.cached_size.set(my_size);
        my_size
    }

    fn write_to_with_cached_sizes(&self, os: &mut ::protobuf::CodedOutputStream<'_>) -> ::protobuf::ProtobufResult<()> {
        if self.code != 0 {
            os.write_int32(1, self.code)?;
        }
        if !self.msg.is_empty() {
            os.write_string(2, &self.msg)?;
        }
        os.write_unknown_fields(self.get_unknown_fields())?;
        ::std::result::Result::Ok(())
    }

    fn get_cached_size(&self) -> u32 {
        self.cached_size.get()
    }

    fn get_unknown_fields(&self) -> &::protobuf::UnknownFields {
        &self.unknown_fields
    }

    fn mut_unknown_fields(&mut self) -> &mut ::protobuf::UnknownFields {
        &mut self.unknown_fields
    }

    fn as_any(&self) -> &dyn (::std::any::Any) {
        self as &dyn (::std::any::Any)
    }
    fn as_any_mut(&mut self) -> &mut dyn (::std::any::Any) {
        self as &mut dyn (::std::any::Any)
    }
    fn into_any(self: ::std::boxed::Box<Self>) -> ::std::boxed::Box<dyn (::std::any::Any)> {
        self
    }

    fn descriptor(&self) -> &'static ::protobuf::reflect::MessageDescriptor {
        Self::descriptor_static()
    }

    fn new() -> EventResponse {
        EventResponse::new()
    }

    fn descriptor_static() -> &'static ::protobuf::reflect::MessageDescriptor {
        static descriptor: ::protobuf::rt::LazyV2<::protobuf::reflect::MessageDescriptor> = ::protobuf::rt::LazyV2::INIT;
        descriptor.get(|| {
            let mut fields = ::std::vec::Vec::new();
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeInt32>(
                "code",
                |m: &EventResponse| { &m.code },
                |m: &mut EventResponse| { &mut m.code },
            ));
            fields.push(::protobuf::reflect::accessor::make_simple_field_accessor::<_, ::protobuf::types::ProtobufTypeString>(
                "msg",
                |m: &EventResponse| { &m.msg },
                |m: &mut EventResponse| { &mut m.msg },
            ));
            ::protobuf::reflect::MessageDescriptor::new_pb_name::<EventResponse>(
                "EventResponse",
                fields,
                file_descriptor_proto()
            )
        })
    }

    fn default_instance() -> &'static EventResponse {
        static instance: ::protobuf::rt::LazyV2<EventResponse> = ::protobuf::rt::LazyV2::INIT;
        instance.get(EventResponse::new)
    }
}

impl ::protobuf::Clear for EventResponse {
    fn clear(&mut self) {
        self.code = 0;
        self.msg.clear();
        self.unknown_fields.clear();
    }
}

impl ::std::fmt::Debug for EventResponse {
    fn fmt(&self, f: &mut ::std::fmt::Formatter<'_>) -> ::std::fmt::Result {
        ::protobuf::text_format::fmt(self, f)
    }
}

impl ::protobuf::reflect::ProtobufValue for EventResponse {
    fn as_ref(&self) -> ::protobuf::reflect::ReflectValueRef {
        ::protobuf::reflect::ReflectValueRef::Message(self)
    }
}

static file_descriptor_proto_data: &'static [u8] = b"\
    \n\x12common/probe.proto\x12\x06common\"\xed\x01\n\x0cProbeRequest\x129\
    \n\x08nodeType\x18\x01\x20\x01(\x0e2\x1d.common.ProbeRequest.NodeTypeR\
    \x08nodeType\x12<\n\tprobeType\x18\x02\x20\x01(\x0e2\x1e.common.ProbeReq\
    uest.ProbeTypeR\tprobeType\":\n\x08NodeType\x12\x0f\n\x0bCoordinator\x10\
    \0\x12\x0e\n\nTaskWorker\x10\x01\x12\r\n\tConnector\x10\x02\"(\n\tProbeT\
    ype\x12\x0c\n\x08Liveness\x10\0\x12\r\n\tReadiness\x10\x01\"W\n\rProbeRe\
    sponse\x12\x16\n\x06memory\x18\x01\x20\x01(\x02R\x06memory\x12\x10\n\x03\
    cpu\x18\x02\x20\x01(\x02R\x03cpu\x12\x1c\n\tavailable\x18\x03\x20\x01(\
    \x08R\tavailable\"\"\n\x0cEventRequest\x12\x12\n\x04data\x18\x01\x20\x01\
    (\x0cR\x04data\"5\n\rEventResponse\x12\x12\n\x04code\x18\x01\x20\x01(\
    \x05R\x04code\x12\x10\n\x03msg\x18\x02\x20\x01(\tR\x03msgB\x1eZ\x1ctable\
    flow/alpha/common/probeJ\xec\x06\n\x06\x12\x04\0\0\x20\x01\n\x08\n\x01\
    \x0c\x12\x03\0\0\x12\n\x08\n\x01\x02\x12\x03\x02\0\x0f\n\x08\n\x01\x08\
    \x12\x03\x03\03\n\t\n\x02\x08\x0b\x12\x03\x03\03\n\n\n\x02\x04\0\x12\x04\
    \x05\0\x11\x01\n\n\n\x03\x04\0\x01\x12\x03\x05\x08\x14\n\x0c\n\x04\x04\0\
    \x04\0\x12\x04\x06\x02\n\x03\n\x0c\n\x05\x04\0\x04\0\x01\x12\x03\x06\x07\
    \x0f\n\r\n\x06\x04\0\x04\0\x02\0\x12\x03\x07\x04\x14\n\x0e\n\x07\x04\0\
    \x04\0\x02\0\x01\x12\x03\x07\x04\x0f\n\x0e\n\x07\x04\0\x04\0\x02\0\x02\
    \x12\x03\x07\x12\x13\n\r\n\x06\x04\0\x04\0\x02\x01\x12\x03\x08\x04\x13\n\
    \x0e\n\x07\x04\0\x04\0\x02\x01\x01\x12\x03\x08\x04\x0e\n\x0e\n\x07\x04\0\
    \x04\0\x02\x01\x02\x12\x03\x08\x11\x12\n\r\n\x06\x04\0\x04\0\x02\x02\x12\
    \x03\t\x04\x12\n\x0e\n\x07\x04\0\x04\0\x02\x02\x01\x12\x03\t\x04\r\n\x0e\
    \n\x07\x04\0\x04\0\x02\x02\x02\x12\x03\t\x10\x11\n\x0c\n\x04\x04\0\x04\
    \x01\x12\x04\x0b\x02\x0e\x03\n\x0c\n\x05\x04\0\x04\x01\x01\x12\x03\x0b\
    \x07\x10\n\r\n\x06\x04\0\x04\x01\x02\0\x12\x03\x0c\x04\x11\n\x0e\n\x07\
    \x04\0\x04\x01\x02\0\x01\x12\x03\x0c\x04\x0c\n\x0e\n\x07\x04\0\x04\x01\
    \x02\0\x02\x12\x03\x0c\x0f\x10\n\r\n\x06\x04\0\x04\x01\x02\x01\x12\x03\r\
    \x04\x12\n\x0e\n\x07\x04\0\x04\x01\x02\x01\x01\x12\x03\r\x04\r\n\x0e\n\
    \x07\x04\0\x04\x01\x02\x01\x02\x12\x03\r\x10\x11\n\x0b\n\x04\x04\0\x02\0\
    \x12\x03\x0f\x02\x18\n\x0c\n\x05\x04\0\x02\0\x06\x12\x03\x0f\x02\n\n\x0c\
    \n\x05\x04\0\x02\0\x01\x12\x03\x0f\x0b\x13\n\x0c\n\x05\x04\0\x02\0\x03\
    \x12\x03\x0f\x16\x17\n\x0b\n\x04\x04\0\x02\x01\x12\x03\x10\x02\x1a\n\x0c\
    \n\x05\x04\0\x02\x01\x06\x12\x03\x10\x02\x0b\n\x0c\n\x05\x04\0\x02\x01\
    \x01\x12\x03\x10\x0c\x15\n\x0c\n\x05\x04\0\x02\x01\x03\x12\x03\x10\x18\
    \x19\n\n\n\x02\x04\x01\x12\x04\x13\0\x17\x01\n\n\n\x03\x04\x01\x01\x12\
    \x03\x13\x08\x15\n\x0b\n\x04\x04\x01\x02\0\x12\x03\x14\x02\x13\n\x0c\n\
    \x05\x04\x01\x02\0\x05\x12\x03\x14\x02\x07\n\x0c\n\x05\x04\x01\x02\0\x01\
    \x12\x03\x14\x08\x0e\n\x0c\n\x05\x04\x01\x02\0\x03\x12\x03\x14\x11\x12\n\
    \x0b\n\x04\x04\x01\x02\x01\x12\x03\x15\x02\x10\n\x0c\n\x05\x04\x01\x02\
    \x01\x05\x12\x03\x15\x02\x07\n\x0c\n\x05\x04\x01\x02\x01\x01\x12\x03\x15\
    \x08\x0b\n\x0c\n\x05\x04\x01\x02\x01\x03\x12\x03\x15\x0e\x0f\n\x0b\n\x04\
    \x04\x01\x02\x02\x12\x03\x16\x02\x15\n\x0c\n\x05\x04\x01\x02\x02\x05\x12\
    \x03\x16\x02\x06\n\x0c\n\x05\x04\x01\x02\x02\x01\x12\x03\x16\x07\x10\n\
    \x0c\n\x05\x04\x01\x02\x02\x03\x12\x03\x16\x13\x14\n\n\n\x02\x04\x02\x12\
    \x04\x19\0\x1b\x01\n\n\n\x03\x04\x02\x01\x12\x03\x19\x08\x14\n\x0b\n\x04\
    \x04\x02\x02\0\x12\x03\x1a\x02\x11\n\x0c\n\x05\x04\x02\x02\0\x05\x12\x03\
    \x1a\x02\x07\n\x0c\n\x05\x04\x02\x02\0\x01\x12\x03\x1a\x08\x0c\n\x0c\n\
    \x05\x04\x02\x02\0\x03\x12\x03\x1a\x0f\x10\n\n\n\x02\x04\x03\x12\x04\x1d\
    \0\x20\x01\n\n\n\x03\x04\x03\x01\x12\x03\x1d\x08\x15\n\x0b\n\x04\x04\x03\
    \x02\0\x12\x03\x1e\x02\x11\n\x0c\n\x05\x04\x03\x02\0\x05\x12\x03\x1e\x02\
    \x07\n\x0c\n\x05\x04\x03\x02\0\x01\x12\x03\x1e\x08\x0c\n\x0c\n\x05\x04\
    \x03\x02\0\x03\x12\x03\x1e\x0f\x10\n\x0b\n\x04\x04\x03\x02\x01\x12\x03\
    \x1f\x02\x11\n\x0c\n\x05\x04\x03\x02\x01\x05\x12\x03\x1f\x02\x08\n\x0c\n\
    \x05\x04\x03\x02\x01\x01\x12\x03\x1f\t\x0c\n\x0c\n\x05\x04\x03\x02\x01\
    \x03\x12\x03\x1f\x0f\x10b\x06proto3\
";

static file_descriptor_proto_lazy: ::protobuf::rt::LazyV2<::protobuf::descriptor::FileDescriptorProto> = ::protobuf::rt::LazyV2::INIT;

fn parse_descriptor_proto() -> ::protobuf::descriptor::FileDescriptorProto {
    ::protobuf::Message::parse_from_bytes(file_descriptor_proto_data).unwrap()
}

pub fn file_descriptor_proto() -> &'static ::protobuf::descriptor::FileDescriptorProto {
    file_descriptor_proto_lazy.get(|| {
        parse_descriptor_proto()
    })
}
