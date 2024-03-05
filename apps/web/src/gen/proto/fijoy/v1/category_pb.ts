// @generated by protoc-gen-es v1.7.2 with parameter "target=ts"
// @generated from file fijoy/v1/category.proto (package fijoy.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";
import { TransactionType } from "./transaction_pb.js";

/**
 * @generated from message fijoy.v1.Category
 */
export class Category extends Message<Category> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string workspace_id = 2;
   */
  workspaceId = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: fijoy.v1.TransactionType category_type = 4;
   */
  categoryType = TransactionType.EXPENSE;

  constructor(data?: PartialMessage<Category>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "fijoy.v1.Category";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "workspace_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "category_type", kind: "enum", T: proto3.getEnumType(TransactionType) },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Category {
    return new Category().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Category {
    return new Category().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Category {
    return new Category().fromJsonString(jsonString, options);
  }

  static equals(a: Category | PlainMessage<Category> | undefined, b: Category | PlainMessage<Category> | undefined): boolean {
    return proto3.util.equals(Category, a, b);
  }
}
