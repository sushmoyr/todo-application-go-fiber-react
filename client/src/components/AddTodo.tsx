import { useForm } from '@mantine/form';
import {Button, Group, Modal, Textarea, TextInput} from "@mantine/core";
import {useDisclosure} from "@mantine/hooks";
import {ENDPOINT, Todo} from "../App.tsx";
import {KeyedMutator} from "swr";

interface TodoForm {
    title: string,
    body: string
}

function AddTodo({mutate} : { mutate: KeyedMutator<Todo[]> }) {
    const [opened, { open, close }] = useDisclosure(false);

    const form = useForm({
        initialValues: {
            title: "",
            body: "",
        },
    });

    async function createTodo(value: TodoForm) {
        const updated = await fetch(`${ENDPOINT}/api/todos`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify(value),
        }).then((r) => r.json());

        mutate(updated)
        form.reset();
        close();
    }

    return (
        <>
            <Modal opened={opened} onClose={close} title="Create Todo" centered withCloseButton={true}>
                <form onSubmit={form.onSubmit(createTodo)}>
                    <TextInput
                        required
                        mb={12}
                        label={"Todo"}
                        placeholder={"What do you want to do?"}
                        {...form.getInputProps("title")}
                    />
                    <Textarea
                        required
                        mb={12}
                        label={"Description"}
                        placeholder={"What do you want to do?"}
                        {...form.getInputProps("body")}
                    />

                    <Button type={"submit"}>Create TODO</Button>
                </form>
            </Modal>

            <Group position="center">
                <Button fullWidth onClick={open}>Open centered Modal</Button>
            </Group>
        </>
    );
}

export default AddTodo;
