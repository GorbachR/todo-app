import { useState } from "react";
import TodoContainer from "./components/TodoContainer/TodoContainer";
import ThemeContext from "./context/ThemeContext";
import { QueryClient, QueryClientProvider } from "react-query";
import "./App.css";

const queryClient = new QueryClient();

export default function App() {
  const [theme, setTheme] = useState("light");

  return (
    <QueryClientProvider client={queryClient}>
      <ThemeContext.Provider value={{ theme, setTheme }}>
        <div className={theme + "-theme" + " min-h-screen bg-todo-bg"}>
          <div className={`h-[300px] bg-cover bg-no-repeat bg-${theme}`}>
            <TodoContainer />
          </div>
        </div>
      </ThemeContext.Provider>
    </QueryClientProvider>
  );
}
