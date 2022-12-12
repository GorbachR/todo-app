import { Routes, Route, Navigate } from "react-router-dom";
import Login from "./Components/Login/Login";
import Home from "./Components/Home/Home";
import NoteLayout from "./Components/NoteLayout/NoteLayout";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";

function App() {
  const client = new QueryClient();
  return (
    <QueryClientProvider client={client}>
      <Routes>
        <Route path="/" element={<Navigate to="/login" />} />
        <Route path="/note" element={<NoteLayout />}>
          <Route index element={<Navigate to="/note/1" />} />
          <Route path=":id" element={<Home />} />
        </Route>
        <Route path="/login" element={<Login />} />
        <Route path="*" element={<Navigate to="/" />} />
      </Routes>
    </QueryClientProvider>
  );
}

export default App;
