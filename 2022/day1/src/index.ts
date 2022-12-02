import { readFileSync } from "fs";

const elfTotals: Array<number> = readFileSync('src/day1-input.txt', 'utf-8')
.toString()
.split(/\n\s*\n/)
.map(elf => elf.split(/\n/))
.map(elfCalorieList => elfCalorieList
  .map(v => +v)
  .reduce((acc, curr) => acc + curr, 0)
  );

const challenge1 = Math.max(...elfTotals);

const challenge2 = elfTotals
.slice(0, 3)
.reduce((acc, curr) => acc + curr, 0);

console.log({ challenge1, challenge2 });