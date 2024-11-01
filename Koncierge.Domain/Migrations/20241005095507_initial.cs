using System;
using Microsoft.EntityFrameworkCore.Migrations;

#nullable disable

namespace Koncierge.Domain.Migrations
{
    /// <inheritdoc />
    public partial class initial : Migration
    {
        /// <inheritdoc />
        protected override void Up(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.CreateTable(
                name: "KubeConfigs",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
                    Name = table.Column<string>(type: "TEXT", nullable: false),
                    Path = table.Column<string>(type: "TEXT", nullable: false),
                    IsDefault = table.Column<bool>(type: "INTEGER", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_KubeConfigs", x => x.Id);
                });

            migrationBuilder.CreateTable(
                name: "Forwards",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
                    Context = table.Column<string>(type: "TEXT", nullable: false),
                    Type = table.Column<int>(type: "INTEGER", nullable: false),
                    Selector = table.Column<string>(type: "TEXT", nullable: false),
                    LocalPort = table.Column<int>(type: "INTEGER", nullable: false),
                    RemotePort = table.Column<int>(type: "INTEGER", nullable: false),
                    WithConfigId = table.Column<Guid>(type: "TEXT", nullable: false),
                    Namespace = table.Column<string>(type: "TEXT", nullable: false)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_Forwards", x => x.Id);
                    table.ForeignKey(
                        name: "FK_Forwards_KubeConfigs_WithConfigId",
                        column: x => x.WithConfigId,
                        principalTable: "KubeConfigs",
                        principalColumn: "Id",
                        onDelete: ReferentialAction.Cascade);
                });

            migrationBuilder.CreateTable(
                name: "LinkedConfigs",
                columns: table => new
                {
                    Id = table.Column<Guid>(type: "TEXT", nullable: false),
                    Type = table.Column<int>(type: "INTEGER", nullable: false),
                    Name = table.Column<string>(type: "TEXT", nullable: false),
                    Values = table.Column<string>(type: "TEXT", nullable: false),
                    ForwardEntityId = table.Column<Guid>(type: "TEXT", nullable: true)
                },
                constraints: table =>
                {
                    table.PrimaryKey("PK_LinkedConfigs", x => x.Id);
                    table.ForeignKey(
                        name: "FK_LinkedConfigs_Forwards_ForwardEntityId",
                        column: x => x.ForwardEntityId,
                        principalTable: "Forwards",
                        principalColumn: "Id");
                });

            migrationBuilder.CreateIndex(
                name: "IX_Forwards_WithConfigId",
                table: "Forwards",
                column: "WithConfigId");

            migrationBuilder.CreateIndex(
                name: "IX_LinkedConfigs_ForwardEntityId",
                table: "LinkedConfigs",
                column: "ForwardEntityId");
        }

        /// <inheritdoc />
        protected override void Down(MigrationBuilder migrationBuilder)
        {
            migrationBuilder.DropTable(
                name: "LinkedConfigs");

            migrationBuilder.DropTable(
                name: "Forwards");

            migrationBuilder.DropTable(
                name: "KubeConfigs");
        }
    }
}
