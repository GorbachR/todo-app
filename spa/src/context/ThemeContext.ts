import { createContext } from "react";

type ThemeContext = {
  theme: string;
  setTheme: React.Dispatch<React.SetStateAction<string>>;
};

const Context = createContext<ThemeContext | null>(null);

export default Context;
