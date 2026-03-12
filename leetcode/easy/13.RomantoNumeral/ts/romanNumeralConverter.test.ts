import { expect, test, describe } from "bun:test";
import romanToInt from "./romanNumeralConverter"; // Assuming your file is named roman.ts

describe("Roman to Integer", () => {
  test("Example 1: Basic addition (III)", () => {
    expect(romanToInt("III")).toBe(3);
  });

  test("Example 2: Subtraction rule (LVIII)", () => {
    expect(romanToInt("LVIII")).toBe(58);
  });

  test("Example 3: Complex subtraction (MCMXCIV)", () => {
    expect(romanToInt("MCMXCIV")).toBe(1994);
  });

  test("Single characters", () => {
    expect(romanToInt("I")).toBe(1);
    expect(romanToInt("M")).toBe(1000);
  });

  test("Boundary cases (IV and IX)", () => {
    expect(romanToInt("IV")).toBe(4);
    expect(romanToInt("IX")).toBe(9);
  });
});