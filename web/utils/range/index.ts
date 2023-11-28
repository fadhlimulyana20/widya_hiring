export const range = (start: number, stop: number, step: number) =>
  Array.from({ length: (stop - start) }, (_, i) => start + i * step);
