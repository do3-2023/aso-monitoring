import { MigrationInterface, QueryRunner } from 'typeorm';

export class V2NullablePersonLocation1718202607833
  implements MigrationInterface
{
  public async up(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE person ALTER COLUMN location DROP NOT NULL;`,
    );
  }

  public async down(queryRunner: QueryRunner): Promise<void> {
    await queryRunner.query(
      `ALTER TABLE person ALTER COLUMN location SET NOT NULL;`,
    );
  }
}
