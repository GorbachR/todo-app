import { useState, useEffect, useRef } from "react";

interface IWindowSize {
  x: Number;
  y: Number;
}

export default function useWindowDimensions(): IWindowSize {
  const [windowSize, setWindowSize] = useState<IWindowSize>({ x: 0, y: 0 });
  let timeout: ReturnType<typeof setTimeout>;

  useEffect(() => {
    let throttle: boolean = false;

    function handleResize(): void {
      if (throttle) return;

      throttle = true;
      setTimeout(() => {
        setWindowSize({ x: window.innerWidth, y: window.innerHeight });
        throttle = false;
      }, 250);
    }

    window.addEventListener("resize", handleResize);

    return () => {
      window.removeEventListener("resize", handleResize);
    };
  }, []);

  return windowSize;
}
