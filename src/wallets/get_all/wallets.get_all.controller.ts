import { Body, Controller, Get } from '@nestjs/common';
import { WalletsGetAllService } from './wallets.get_all.service';

@Controller('wallets')
export class WalletsGetAllController {
  constructor(private readonly walletsGetAllService: WalletsGetAllService) {}

  @Get()
  public async getAll() {
    return await this.walletsGetAllService.execute();
  }
}
