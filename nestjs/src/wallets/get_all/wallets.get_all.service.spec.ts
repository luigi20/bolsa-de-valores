import { Test, TestingModule } from '@nestjs/testing';
import { WalletsGetAllService } from './wallets.get_all.service';

describe('AssetsService', () => {
  let service: WalletsGetAllService;

  beforeEach(async () => {
    const module: TestingModule = await Test.createTestingModule({
      providers: [WalletsGetAllService],
    }).compile();

    service = module.get<WalletsGetAllService>(WalletsGetAllService);
  });

  it('should be defined', () => {
    expect(service).toBeDefined();
  });
});
