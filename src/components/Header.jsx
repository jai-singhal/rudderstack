const Header = ({ head, description }) => {
  return (
      <div className='text-center mt-5'>
        <h3>{head}</h3>
        <p className='text-capitalize'>{description}</p>
      </div>
  )
}

export default Header