import { useEffect, useState, ChangeEvent } from "react";
import { AgGridReact } from "ag-grid-react";
import "ag-grid-community/styles/ag-grid.css";
import "ag-grid-community/styles/ag-theme-alpine.css";
import { ColDef } from "ag-grid-community";
import { ClassDivisions } from "@/services/types";
import ClassSelector from "../components/ClassSelector";
import { Button } from "../components/ui/button";
import { useGetSchedule } from "@/services/mutations";
import { Box, CircularProgress } from "@mui/material";

type ScheduleRow = {
  Mon: string;
  Tue: string;
  Wed: string;
  Thu: string;
  Fri: string;
  Sat: string;
};

function Schedule() {
  const [showLoader, setShowLoader] = useState(false);
  const [classes] = useState<ClassDivisions>({
    ClassNames: [],
    Divisions: {},
  });
  const [selectedClass, setSelectedClass] = useState<string>("");
  const [selectedDivision, setSelectedDivision] = useState<string>("");
  const [avgResponse, setAvgResponse] = useState<string>("50");
  const [rowData, setRowData] = useState<ScheduleRow[]>([]);
  const [colDefs] = useState<ColDef[]>([
    {
      field: "Mon",
      width: 100,
      resizable: false,
      flex: 1,
      sortable: false,
      headerClass: "ag-center-header",
      suppressMovable: true,
    },
    {
      field: "Tue",
      width: 100,
      resizable: false,
      flex: 1,
      sortable: false,
      headerClass: "ag-center-header",
      suppressMovable: true,
    },
    {
      field: "Wed",
      width: 100,
      resizable: false,
      flex: 1,
      sortable: false,
      headerClass: "ag-center-header",
      suppressMovable: true,
    },
    {
      field: "Thu",
      width: 100,
      resizable: false,
      flex: 1,
      sortable: false,
      headerClass: "ag-center-header",
      suppressMovable: true,
    },
    {
      field: "Fri",
      width: 100,
      resizable: false,
      flex: 1,
      sortable: false,
      headerClass: "ag-center-header",
      suppressMovable: true,
    },
    {
      field: "Sat",
      width: 100,
      resizable: false,
      flex: 1,
      sortable: false,
      headerClass: "ag-center-header",
      suppressMovable: true,
    },
  ]);

  // const getClassesQuery = useGetClassInfo();
  const getScheduleMutation = useGetSchedule();

  // useEffect(() => {
  //   if (getClassesQuery.data) {
  //     setClasses(getClassesQuery.data);
  //     const classNames = getClassesQuery.data.ClassNames;
  //     if (classNames.length > 0) {
  //       setSelectedClass(classNames[0]);
  //     }
  //   }
  // }, [getClassesQuery.data]);

  const getScheduleData = async () => {
    setShowLoader(true);
    const payload = {
      model: "llama3",
      prompt: JSON.stringify({
        subjects: [
          { teacher_name: "John Doe", subject_name: "Python", occurrence: 5 },
          { teacher_name: "Jane Smith", subject_name: "SPM", occurrence: 5 },
          {
            teacher_name: "Alice Johns",
            subject_name: "OT",
            occurrence: 6,
          },
          { teacher_name: "Bob Brown", subject_name: "ADBMS", occurrence: 4 },
          {
            teacher_name: "Charlie Green",
            subject_name: "OC3",
            occurrence: 5,
          },
          { teacher_name: "Emily White", subject_name: "OC4", occurrence: 6 },
          {
            teacher_name: "Thomas Edison",
            subject_name: "Java",
            occurrence: 5,
          },
        ],
      }),
    };

    let schedule: any;

    getScheduleMutation.mutate(payload, {
      onSuccess: (data) => {
        schedule = JSON.parse(data.response);

        // Create row data ensuring there are exactly 6 rows
        const maxRows = 6;
        const rows: ScheduleRow[] = Array.from(
          { length: maxRows },
          (_, i) =>
            Object.fromEntries(
              Object.keys(schedule).map((day) => [day, schedule[day][i] || ""])
            ) as ScheduleRow
        );

        setRowData(rows);
        let total_duration = data.total_duration / 1000000000;
        setAvgResponse(total_duration.toString().substring(0, 2));
        setShowLoader(false);
      },
    });
  };

  useEffect(() => {
    const selectedClassDivisions = classes.Divisions[selectedClass] || [];
    setSelectedDivision(selectedClassDivisions[0] || "");
  }, [selectedClass, classes]);

  const handleClassChange = (event: ChangeEvent<HTMLSelectElement>) => {
    const selectedClass = event.target.value;
    setSelectedClass(selectedClass);
    const selectedClassDivisions = classes.Divisions[selectedClass] || [];
    setSelectedDivision(selectedClassDivisions[0] || "");
  };

  const handleDivisionChange = (event: ChangeEvent<HTMLSelectElement>) => {
    setSelectedDivision(event.target.value);
  };

  const handleSearch = () => {
    getScheduleData();
  };

  return (
    <>
      <div className="mt-7">
        <h2 className="font-bold text-2xl flex justify-between items-center text-gray-700">
          Generate Schedule
        </h2>
        <div className="flex justify-between items-center my-3 border rounded-lg shadow-sm p-3 bg-white">
          <div className="flex gap-6">
            <ClassSelector
              selectedClass={selectedClass}
              selectedDivision={selectedDivision}
              classes={classes}
              handleClassChange={handleClassChange}
              handleDivisionChange={handleDivisionChange}
            />
            <Button onClick={handleSearch}>Search</Button>
          </div>
          <div className="text-xs text-gray-500 mr-4">
            Avg. Response: {avgResponse} s
          </div>
        </div>
        <div className="mt-4">
          <div className="ag-theme-alpine h-[530px]">
            {showLoader ? (
              <div className="flex justify-center items-center h-full bg-slate-200 rounded-lg">
                <Box
                  sx={{
                    display: "flex",
                    alignItems: "center",
                    justifyContent: "center",
                  }}
                >
                  <CircularProgress />
                </Box>
              </div>
            ) : (
              <AgGridReact<ScheduleRow>
                rowData={rowData}
                columnDefs={colDefs}
              />
            )}
          </div>
        </div>
      </div>
    </>
  );
}

export default Schedule;
