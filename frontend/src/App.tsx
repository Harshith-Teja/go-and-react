import { Container, Stack } from "@chakra-ui/react";
import Navbar from "./components/Navbar";
import TodoForm from "./components/TodoForm";
import TodoList from "./components/TodoList";

function App() {
  return (
    <Stack h="100vh">
      <Navbar />
      <Container w="40vw">
        <TodoForm />
        <TodoList />
      </Container>
    </Stack>
  );
}

export default App;
