module.exports = {
    devServer: {
        headers: {
          'Cache-Control': 'must-revalidate',
          'Vary': '*',         
        }
      }
  }
  