import "./App.css";
import Home from "./components/Home";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { ThemeProvider } from "@mui/material/styles";

import Dashboard from "./pages/Dashboard";
import Settings from "./pages/Settings";
import theme from "./config/AppTheme";
import Login from "./pages/Login";
import Schedule from "./pages/Schedule";
import Loader from "./components/Loader";

const router = createBrowserRouter([
  {
    path: "/",
    element: <Home />,
    children: [
      { path: "dashboard", element: <Dashboard /> },
      { path: "schedule", element: <Schedule /> },
      { path: "settings", element: <Settings /> },
    ],
  },
  { path: "login", element: <Login /> },
]);
function App() {
  return (
    <ThemeProvider theme={theme}>
      <RouterProvider router={router} />
      <Loader />
    </ThemeProvider>
  );
}

export default App;
