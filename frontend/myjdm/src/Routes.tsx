import {
    BrowserRouter as Router,
    Routes,
    Route,
} from "react-router-dom";
import { PageCars } from "./pages/Cars";
import { PageCar } from "./pages/Car";

export function AppRoutes(){
    return(
        <Router>
            <Routes>
                <Route path="/cars" element={<PageCars />}/>
                <Route path="/car/:id" element={<PageCar />}/>
            </Routes>
        </Router>
    )
}