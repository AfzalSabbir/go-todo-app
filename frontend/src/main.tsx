import {StrictMode} from 'react'
import {createRoot} from 'react-dom/client'
import {RouterProvider} from "react-router-dom";

import {Toaster} from "./components/ui/sonner.tsx";

import "./index.css";
import routes from "./routes/router";

createRoot(document.getElementById('root')!).render(
    <StrictMode>
        <Toaster />
        <RouterProvider router={routes}/>
    </StrictMode>,
)
