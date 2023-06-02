import './App.css'
import {Box} from "@mantine/core";
import useSWR from "swr";


export const ENDPOINT = "http://localhost:4000";
const fetcher = (url: string) => fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

function App() {

    const {data, mutate} = useSWR('api/todos', fetcher);

  return <Box>
          Hello World
      </Box>

}

export default App
