export interface TunableStruct {
  stringField?: string;
  numberField?: number;
  booleanField?: boolean;
}

/** tunable boolean */
export function getBooleanTunable(): boolean {
  return true;
}

/** tunable number */
export function getNumberTunable(): number {
  return 42;
}

/** tunable string */
export function getStringTunable(): string {
  return "foo";
}

/** test boolean operators */
export function getTestBooleanOperators({
  isTest,
}: {
  isTest: boolean;
}): string {
  if (isTest) {
    return "true";
  } else if (!isTest) {
    return "false";
  }
  return "default";
}

/** test number operators */
export function getTestNumberOperators({
  version,
}: {
  version: number;
}): string {
  if (version === 1) {
    return "equals";
  } else if (version !== 2) {
    return "not equals";
  } else if (version > 3) {
    return "greater";
  } else if (version >= 4) {
    return "greater or equals";
  } else if (version < 5) {
    return "less";
  } else if (version <= 6) {
    return "less or equals";
  } else if ([1, 3, 5].includes(version)) {
    return "in";
  }
  return "default";
}

/** test string operators */
export function getTestStringOperators({ env }: { env: string }): string {
  if (env === "prod") {
    return "equals";
  } else if (env !== "dev") {
    return "not equals";
  } else if (env.includes("test")) {
    return "test";
  } else if (env.startsWith("staging")) {
    return "staging";
  } else if (env.endsWith("beta")) {
    return "beta";
  } else if (["special1", "special2"].includes(env)) {
    return "special";
  }
  return "default";
}

/** tunable interface */
export function getTunableInterface(): TunableStruct {
  return {
    booleanField: true,
    numberField: 42,
    stringField: "default",
  };
}
