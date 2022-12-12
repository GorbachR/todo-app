import { useQuery } from "@tanstack/react-query";
import axios from "axios";

export default function useGetNotes() {
  return useQuery(["products"], () => {
    return axios.get(import.meta.env.VITE_NOTESAPI).then((res) => res.data);
  });
}
