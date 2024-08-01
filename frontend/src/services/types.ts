export type ClassDivisions = {
  ClassNames: string[];
  Divisions: {
    [key: string]: string[];
  };
};

export type ScheduleData = {
  model: string;
  prompt: string;
};

export type Message = {
  role: string;
  content: string;
};

export type StudentData = {
  name: string;
}
