
import { Context } from './Context'


class CompanySearchError extends Error {

  isCompanySearchError = true

  sdk = 'CompanySearch'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  CompanySearchError
}

