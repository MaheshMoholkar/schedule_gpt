import { useMutation } from "@tanstack/react-query";
import { toast } from "sonner";
import { getChat, getSchedule } from "./api";
import { Message, ScheduleData } from "./types";

export function useGetSchedule() {
  return useMutation({
    mutationFn: (data: ScheduleData) => getSchedule(data),
    onSuccess: async () => {
      toast("Schedule Fetched!");
    },
  });
}

export function useChat() {
  return useMutation({
    mutationFn: (data: Message[]) => getChat(data),
  });
}


export function useCreateStudent() {
  return useMutation({
    mutationFn: (data: Message[]) => getChat(data),
  });
}

