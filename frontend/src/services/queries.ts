import { useQuery } from "@tanstack/react-query";
import { getClassDivisions } from "./api";

export function useGetClassInfo() {
  return useQuery({
    queryKey: ["class-divisions"],
    queryFn: getClassDivisions,
    refetchOnWindowFocus: false,
  });
}
