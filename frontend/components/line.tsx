interface LineParams {
    lineHeight?: number
    horizontal?: boolean
    strokeWidth?: number
}

export const SVGLine = ({ lineHeight = 100, horizontal = false, strokeWidth = 5} : LineParams) => {
  const [x1, x2, y1, y2]: [number, number, number, number] = createPoints(lineHeight, horizontal, strokeWidth)
  const width: number = horizontal ? lineHeight : strokeWidth
  const height: number = horizontal ? strokeWidth : lineHeight
  return (
    <svg height={height} width={width}>
      <line
        x1={x1}
        y1={y1}
        x2={x2}
        y2={y2}
        className={`stroke-text-color`}
        style={{strokeWidth: `${strokeWidth}px`}}
      />
    </svg>
  );
};

// return the points of a line in tuple format: [x1, x2, y1, y2]
function createPoints(lineHeight: number, horizontal: boolean, strokeWidth: number): [number, number, number, number] {
  const mid = strokeWidth / 2; 
  if(horizontal) {
    return [0, lineHeight, mid, mid]
  } else {
    return [mid, mid, 0, lineHeight]
  }
}