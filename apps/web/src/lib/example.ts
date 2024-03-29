import { createPromiseClient } from "@connectrpc/connect";
import { Workspace } from "@/gen/proto/fijoy/v1/workspace_pb";
import { AccountService } from "@/gen/proto/fijoy/v1/account_connect";
import { finalTransport } from "./connect";
import { AccountType } from "@/gen/proto/fijoy/v1/account_pb";
import { TransactionService } from "@/gen/proto/fijoy/v1/transaction_connect";
import { Timestamp } from "@bufbuild/protobuf";

const accountClient = createPromiseClient(AccountService, finalTransport);
const transactionClient = createPromiseClient(
  TransactionService,
  finalTransport,
);

export async function populateExample(workspace: Workspace) {
  const cobalt = await accountClient.createAccount(
    {
      name: "Cobalt",
      accountType: AccountType.DEBT,
      currency: "CAD",
      institution: "Amex",
    },
    {
      headers: {
        "Fijoy-Workspace-Id": workspace.id,
      },
    },
  );
  await transactionClient.createAdjustmentTransaction(
    {
      accountId: cobalt.id,
      amount: {
        units: -BigInt(2333),
        nanos: -690000000,
      },
      datetime: Timestamp.now(),
    },
    {
      headers: {
        "Fijoy-Workspace-Id": workspace.id,
      },
    },
  );

  const wealthsimpleCash = await accountClient.createAccount(
    {
      name: "Wealthsimple Cash",
      accountType: AccountType.CASH,
      currency: "CAD",
      institution: "Wealthsimple",
    },
    {
      headers: {
        "Fijoy-Workspace-Id": workspace.id,
      },
    },
  );
  await transactionClient.createAdjustmentTransaction(
    {
      accountId: wealthsimpleCash.id,
      amount: {
        units: BigInt(69420),
        nanos: 690000000,
      },
      datetime: Timestamp.now(),
    },
    {
      headers: {
        "Fijoy-Workspace-Id": workspace.id,
      },
    },
  );

  const tangerineSavings = await accountClient.createAccount(
    {
      name: "Tangerine Savings",
      accountType: AccountType.CASH,
      currency: "CAD",
      institution: "Tangerine",
    },
    {
      headers: {
        "Fijoy-Workspace-Id": workspace.id,
      },
    },
  );
  await transactionClient.createAdjustmentTransaction(
    {
      accountId: tangerineSavings.id,
      amount: {
        units: BigInt(9999),
        nanos: 990000000,
      },
      datetime: Timestamp.now(),
    },
    {
      headers: {
        "Fijoy-Workspace-Id": workspace.id,
      },
    },
  );

  const tangerineChequing = await accountClient.createAccount(
    {
      name: "Tangerine Chequing",
      accountType: AccountType.CASH,
      currency: "CAD",
      institution: "Tangerine",
    },
    {
      headers: {
        "Fijoy-Workspace-Id": workspace.id,
      },
    },
  );
  await transactionClient.createAdjustmentTransaction(
    {
      accountId: tangerineChequing.id,
      amount: {
        units: BigInt(6666),
        nanos: 660000000,
      },
      datetime: Timestamp.now(),
    },
    {
      headers: {
        "Fijoy-Workspace-Id": workspace.id,
      },
    },
  );
}
