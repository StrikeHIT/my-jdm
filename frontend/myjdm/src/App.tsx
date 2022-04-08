import { AppRoutes } from "./Routes"
import { Footer } from "./components/Footer"
import { Header } from "./components/Header"

function App() {
  return (
    <div className="App">
      <Header />
        <AppRoutes></AppRoutes>
      <Footer />
    </div>
  )
}

export default App
