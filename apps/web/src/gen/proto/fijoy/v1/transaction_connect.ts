// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts"
// @generated from file fijoy/v1/transaction.proto (package fijoy.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { CreateAdjustmentTransactionRequest, CreateIncomeTransactionRequest, Transaction, Transactions } from "./transaction_pb.js";
import { Empty, MethodIdempotency, MethodKind } from "@bufbuild/protobuf";

/**
 * @generated from service fijoy.v1.TransactionService
 */
export const TransactionService = {
  typeName: "fijoy.v1.TransactionService",
  methods: {
    /**
     * @generated from rpc fijoy.v1.TransactionService.CreateIncomeTransaction
     */
    createIncomeTransaction: {
      name: "CreateIncomeTransaction",
      I: CreateIncomeTransactionRequest,
      O: Transaction,
      kind: MethodKind.Unary,
    },
    /**
     * rpc CreateExpenseTransaction(CreateExpenseTransactionRequest) returns (Transaction);
     * rpc CreateTransferTransaction(CreateTransferTransactionRequest) returns (Transaction);
     *
     * @generated from rpc fijoy.v1.TransactionService.CreateAdjustmentTransaction
     */
    createAdjustmentTransaction: {
      name: "CreateAdjustmentTransaction",
      I: CreateAdjustmentTransactionRequest,
      O: Transaction,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc fijoy.v1.TransactionService.GetTransactions
     */
    getTransactions: {
      name: "GetTransactions",
      I: Empty,
      O: Transactions,
      kind: MethodKind.Unary,
      idempotency: MethodIdempotency.NoSideEffects,
    },
  }
} as const;

