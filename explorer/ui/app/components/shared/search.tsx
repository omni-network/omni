export function SearchBar() {
  return (
    <div className="form-control">
      <input
        type="text"
        placeholder="Search for block, txn or address"
        className="input input-bordered input-primary w-full max-w-xs ml-3 mr-3 m-1"
      />
    </div>
  )
}
