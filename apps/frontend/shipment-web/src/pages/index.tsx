import useLocalStorage from "react-use-localstorage"

const Page = () => {

    const [token, setToken] = useLocalStorage('jwt', undefined)

    return (
        <h1>Home Page</h1>
    )
}

export default Page