// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file fijoy/v1/workspace.proto (package fijoy.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message fijoy.v1.Workspace
 */
export class Workspace extends Message<Workspace> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string namespace = 2;
   */
  namespace = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: google.protobuf.Timestamp created_at = 4;
   */
  createdAt?: Timestamp;

  /**
   * @generated from field: string primary_currency = 5;
   */
  primaryCurrency = "";

  /**
   * @generated from field: string locale = 6;
   */
  locale = "";

  constructor(data?: PartialMessage<Workspace>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fijoy.v1.Workspace";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "created_at", kind: "message", T: Timestamp },
    { no: 5, name: "primary_currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "locale", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Workspace {
    return new Workspace().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Workspace {
    return new Workspace().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Workspace {
    return new Workspace().fromJsonString(jsonString, options);
  }

  static equals(a: Workspace | PlainMessage<Workspace> | undefined, b: Workspace | PlainMessage<Workspace> | undefined): boolean {
    return proto3.util.equals(Workspace, a, b);
  }
}

/**
 * @generated from message fijoy.v1.Workspaces
 */
export class Workspaces extends Message<Workspaces> {
  /**
   * @generated from field: repeated fijoy.v1.Workspace workspaces = 1;
   */
  workspaces: Workspace[] = [];

  constructor(data?: PartialMessage<Workspaces>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fijoy.v1.Workspaces";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "workspaces", kind: "message", T: Workspace, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Workspaces {
    return new Workspaces().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Workspaces {
    return new Workspaces().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Workspaces {
    return new Workspaces().fromJsonString(jsonString, options);
  }

  static equals(a: Workspaces | PlainMessage<Workspaces> | undefined, b: Workspaces | PlainMessage<Workspaces> | undefined): boolean {
    return proto3.util.equals(Workspaces, a, b);
  }
}

/**
 * @generated from message fijoy.v1.CreateWorkspaceRequest
 */
export class CreateWorkspaceRequest extends Message<CreateWorkspaceRequest> {
  /**
   * @generated from field: string namespace = 1;
   */
  namespace = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string primary_currency = 3;
   */
  primaryCurrency = "";

  /**
   * @generated from field: string locale = 4;
   */
  locale = "";

  constructor(data?: PartialMessage<CreateWorkspaceRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fijoy.v1.CreateWorkspaceRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "primary_currency", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "locale", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateWorkspaceRequest {
    return new CreateWorkspaceRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateWorkspaceRequest {
    return new CreateWorkspaceRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateWorkspaceRequest {
    return new CreateWorkspaceRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateWorkspaceRequest | PlainMessage<CreateWorkspaceRequest> | undefined, b: CreateWorkspaceRequest | PlainMessage<CreateWorkspaceRequest> | undefined): boolean {
    return proto3.util.equals(CreateWorkspaceRequest, a, b);
  }
}

/**
 * @generated from message fijoy.v1.GetWorkspaceByIdRequest
 */
export class GetWorkspaceByIdRequest extends Message<GetWorkspaceByIdRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<GetWorkspaceByIdRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fijoy.v1.GetWorkspaceByIdRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetWorkspaceByIdRequest {
    return new GetWorkspaceByIdRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetWorkspaceByIdRequest {
    return new GetWorkspaceByIdRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetWorkspaceByIdRequest {
    return new GetWorkspaceByIdRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetWorkspaceByIdRequest | PlainMessage<GetWorkspaceByIdRequest> | undefined, b: GetWorkspaceByIdRequest | PlainMessage<GetWorkspaceByIdRequest> | undefined): boolean {
    return proto3.util.equals(GetWorkspaceByIdRequest, a, b);
  }
}

/**
 * @generated from message fijoy.v1.GetWorkspaceByNamespaceRequest
 */
export class GetWorkspaceByNamespaceRequest extends Message<GetWorkspaceByNamespaceRequest> {
  /**
   * @generated from field: string namespace = 1;
   */
  namespace = "";

  constructor(data?: PartialMessage<GetWorkspaceByNamespaceRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fijoy.v1.GetWorkspaceByNamespaceRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "namespace", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): GetWorkspaceByNamespaceRequest {
    return new GetWorkspaceByNamespaceRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): GetWorkspaceByNamespaceRequest {
    return new GetWorkspaceByNamespaceRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): GetWorkspaceByNamespaceRequest {
    return new GetWorkspaceByNamespaceRequest().fromJsonString(jsonString, options);
  }

  static equals(a: GetWorkspaceByNamespaceRequest | PlainMessage<GetWorkspaceByNamespaceRequest> | undefined, b: GetWorkspaceByNamespaceRequest | PlainMessage<GetWorkspaceByNamespaceRequest> | undefined): boolean {
    return proto3.util.equals(GetWorkspaceByNamespaceRequest, a, b);
  }
}

