// @generated by protoc-gen-connect-query v1.3.0 with parameter "target=ts"
// @generated from file proto/fijoy/v1/hello.proto (package fijoy.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { MethodKind } from "@bufbuild/protobuf";
import { HelloRequest, HelloResponse } from "./hello_pb.js";

/**
 * @generated from rpc fijoy.v1.GreetService.Hello
 */
export const hello = {
  localName: "hello",
  name: "Hello",
  kind: MethodKind.Unary,
  I: HelloRequest,
  O: HelloResponse,
  service: {
    typeName: "fijoy.v1.GreetService"
  }
} as const;