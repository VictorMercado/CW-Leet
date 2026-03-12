import { describe, expect, test } from "bun:test";
import highAndLow from "./highAndLow.ts";

describe("Highest and Lowest", () => {
  test("Basic sequence", () => {
    expect(highAndLow("1 2 3 4 5")).toBe("5 1");
  });

  test("Sequence with negative numbers", () => {
    expect(highAndLow("1 2 -3 4 5")).toBe("5 -3");
    expect(highAndLow("1 9 3 4 -5")).toBe("9 -5");
  });

  test("Single number input", () => {
    expect(highAndLow("42")).toBe("42 42");
  });

  test("Identical numbers", () => {
    expect(highAndLow("5 5 5")).toBe("5 5");
  });

  test("Large range", () => {
    expect(highAndLow("10 -10 0")).toBe("10 -10");
  });
});