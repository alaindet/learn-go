import css from './app.module.css';

export function App() {

  return (
    <div className="container">
      <div className="row">

        <div className="col">
          <h1>Movies App</h1>
        </div>

        <div className="col text-end">
          <a href="#" className="btn btn-lg btn-link">
            Login
          </a>
        </div>

        <hr className="mb-3" />

      </div>

      <div className="row">
        <div className="col-md-2">
          <nav>
            <div className="list-group">
              <a href="#" className="list-group-item list-group-item-action">Home</a>
              <a href="#" className="list-group-item list-group-item-action">Movies</a>
            </div>
          </nav>
        </div>
      </div>

    </div>
  )
}
