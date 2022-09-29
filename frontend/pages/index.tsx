import {useState, useEffect} from 'react'

export default function Home() {
  const [data, setData] = useState(null);
  const [error, setError] = useState(null);

  useEffect(() => {
    getData()
  }, []);

  const getData = async () => {
    try {
      const res = await fetch(process.env.NEXT_PUBLIC_BACKEND_URL);
      const data = await res.json();
      setData(data);
    }
    catch (error) {
      setError(error);
    }
  }

  const deleteTask = async (index) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/delete/${index}`, {
        method: 'DELETE'
      });
      const data = await res.json();
      setData(data.message);
      getData()
    }
    catch (error) {
      setError(error);
    }
  }

  const doneTask = async (index) => {
    try {
      const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/changeDone/${index}`, {
        method: 'PUT'
      });
      const data = await res.json();
      setData(data.message);
      getData()
    }
    catch (error) {
      setError(error);
    }
  }

  return (
    <div className='container mx-auto my-5'>
		{error && <div>Failed to load {error.toString()}</div>}
      {
        !data ? <div>Loading...</div>
          : (
            (data?.data ?? []).length === 0 && <p>data kosong</p>
          )
      }

      <Input onSuccess={getData} />
      <table className='table-auto mx-auto mt-5'>
        <thead>
          <tr>
            <th>Check</th>
            <th>Task</th>
            <th>Action</th>
          </tr>
        </thead>
        <tbody>
          {data?.data && data?.data?.map((item, index) => (
            <tr className='my-3' key={index}>
              <td>
                <input className='mx-2' type="checkbox" defaultChecked={item.done} disabled/>
              </td>
              <td>
                <span className='text-rose-300'>ID: {item.ID} task: {item.task}</span>
              </td>
              <td>
                <button className='bg-teal-600
                  hover:bg-teal-700
                  active:bg-teal-700
                  focus:outline-none
                  focus:ring
                  focus:ring-teal-400
                  rounded-md
                  ml-3'
                  onClick={() => doneTask(item.ID)}>Done</button>
                <button className='bg-rose-600
                  hover:bg-rose-700
                  active:bg-rose-700
                  focus:outline-none
                  focus:ring
                  focus:ring-rose-400
                  rounded-md
                  ml-3'
                  onClick={() => deleteTask(item.ID)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
    </div>
  )

  function Input({onSuccess}) {
    const [data, setData] = useState(null);
    const [error, setError] = useState(null);

    const handleSubmit = async (e) => {
      e.preventDefault();
      const formData = new FormData(e.currentTarget);
      const body = {
        task: formData.get("data")
      }
  
      try {
        const res = await fetch(`${process.env.NEXT_PUBLIC_BACKEND_URL}/send`, {
          method: 'POST',
          body: JSON.stringify(body)
        });
        const data = await res.json();
        setData(data.message);
        onSuccess();
      }
      catch (error) {
        setError(error);
      }
    }

    return (
      <div className='text-center'>
        <form onSubmit={handleSubmit}>
          <input className='rounded-md' name="data" type="text" />
          <button className='bg-sky-500
            hover:bg-sky-700
            active:bg-sky-700
            focus:outline-none
            focus:ring
            focus:ring-sky-300
            rounded-md
            ml-3'>Submit</button>
        </form>
      </div>
    )
  }

}
