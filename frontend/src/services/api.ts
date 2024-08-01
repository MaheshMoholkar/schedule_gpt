import axios from "axios";
import { Message, ScheduleData } from "./types";

const BASE_URL = "http://localhost:8003";
const axiosInstance = axios.create({ baseURL: BASE_URL });

export const getClassDivisions = async () => {
  // return (await axiosInstance.get("api/v1/class-divisions")).data;
  return {};
};

export const getSubjects = async () => {
  return (await axiosInstance.get("api/v1/subject")).data;
};

export const getSchedule = async (data: ScheduleData) => {
  return (await axiosInstance.post("api/v1/schedule", data)).data;
};

export const getChat = async (data: Message[]) => {
  return (await axiosInstance.post("api/v1/chat", { messages: data })).data;
};
