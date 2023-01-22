import dayjs from "dayjs"

export const getPriorityClassName = (
  priority: number,
  taskPriority: number
): string => {
  if (priority == taskPriority) {
    return "bg-indigo-700 text-white";
  } else {
    return "bg-gray-200 text-indigo-700";
  }
};

export const getStatusClassName = (taskStatus: string): string => {
  if (taskStatus == "TODO") {
    return "green";
  } else if (taskStatus == "DOING") {
    return "blue";
  } else if (taskStatus == "DONE") {
    return "gray";
  } else {
    return "gray";
  }
};

export const formatDateTime = (date: string): string => {
  return dayjs(date).format("YYYY-MM-DD");
}
