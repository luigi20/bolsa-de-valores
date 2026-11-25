import { Body, Controller, Post } from '@nestjs/common';
import { WalletsCreateService } from './wallets.create.service';

@Controller('wallets')
export class WalletsCreateController {
  constructor(private readonly walletsCreateService: WalletsCreateService) {}

  @Post()
  public async create(@Body() body: { id: string }) {
    return await this.walletsCreateService.execute(body);
  }
}
