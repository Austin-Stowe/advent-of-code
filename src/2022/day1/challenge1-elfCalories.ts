import { readFileSync } from 'fs';

const answer = Math.max(...readFileSync('src/2022/day1/day1-input.txt', 'utf-8')
  .toString()
  .split(/\n\s*\n/)
  .map(elf => elf.split(/\n/))
  .map(elfCalorieList => elfCalorieList
    .map(v => +v)
    .reduce((acc, curr) => acc + curr, 0)
    ));
console.log(answer);