using Koncierge.Domain.Entities;
using Microsoft.EntityFrameworkCore;
using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;
using System.Threading.Tasks;
using static System.Environment;

namespace Koncierge.Domain
{
    public class KonciergeContext : DbContext
    {

        // remove-migration -StartUpProject Koncierge.Domain
        // add-migration <Nome> -StartUpProject Koncierge.Domain -o Migrations
        // add-migration InitialMigration -StartUpProject Koncierge.Domain -o Migrations


        public DbSet<KubeConfigEntity> KubeConfigs { get; set; }
        public DbSet<ForwardEntity> Forwards { get; set; }
        public DbSet<LinkedConfig> LinkedConfigs { get; set; }



        public KonciergeContext() : base()
        {

            //Initialize();

        }
        public KonciergeContext(DbContextOptions<KonciergeContext> options) : base(options)
        {
            Initialize();
        }


        protected override void OnConfiguring(DbContextOptionsBuilder optionsBuilder)
        {
            string AppPath = Path.Combine(Environment.GetFolderPath(SpecialFolder.LocalApplicationData, SpecialFolderOption.DoNotVerify), "koncierge");

            Directory.CreateDirectory(AppPath);


            optionsBuilder.UseSqlite($"Data Source={Path.Combine(AppPath, "koncierge.db")}");



        }



        protected override void OnModelCreating(ModelBuilder modelBuilder)
        {




        }

        public void Initialize()
        {
            Database.Migrate();
           
        }
    }
}
