// import Home from "./pages/Home";
import { Container } from "react-bootstrap";
import Home from "./pages/Home";

function App() {
  return (
    <Container >
      <h1 className='text-center mt-5'>Rudderstack</h1>
      <hr />
      <Home/>
    </Container>
  );
}

export default App;
