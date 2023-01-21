export const getPriorityClass = (
  priority: number,
  taskPriority: number
): string => {
  if (priority == taskPriority) {
    return "bg-indigo-700 text-white";
  } else {
    return "bg-gray-200 text-indigo-700";
  }
};

export const getStatusClass = (taskStatus: string): string => {
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
