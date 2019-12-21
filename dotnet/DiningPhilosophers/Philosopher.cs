using System;
using System.Collections.Generic;
using System.Text;
using System.Threading;
using System.Threading.Tasks;

namespace DiningPhilosophers
{
    class Philosopher
    {
        public string Name { get; set; }

        public Fork LeftFork { get; set; }

        public Fork RightFork { get; set; }

        public Philosopher LeftPhilosopher { get; set; }

        public Philosopher RightPhilosopher { get; set; }

        public Waiter Waiter { get; set; }

        public int ThinkTime { get; set; }

        public int EatTime { get; set; }

        public int ThinkVariance { get; set; }

        public int EatVariance { get; set; }

        public int NextThinkTime => ThinkTime + (int)(-ThinkVariance + new Random().NextDouble() * 2 * ThinkVariance);

        public int NextEatTime => EatTime + (int)(-EatVariance + new Random().NextDouble() * 2 * EatVariance);

        public bool CanEat => !LeftFork.InUse && !RightFork.InUse;

        public void Think()
        {
            Thread.Sleep(NextThinkTime);
        }

        public async Task<bool> Eat()
        {
            //Console.WriteLine(Name + " Requested");
            await Waiter.Request(this);
            LeftFork.Take();
            //Console.WriteLine(Name + " picked up left");
            RightFork.Take();
            //EventManager.Broadcast("Eating");
            //Console.WriteLine(Name + " picked up right");
            Console.WriteLine(Name + " began eating");
            Thread.Sleep(NextEatTime);
            Console.WriteLine(Name + " finished eating");
            LeftFork.Release();
            //Console.WriteLine(Name + " put down left");
            RightFork.Release();
            //Console.WriteLine(Name + " put down right");
            EventManager.Broadcast("Finished", Name);
            return true;
        }

        public async void Run()
        {
            while (true)
            {
                await Eat();
                Think();
            }
        }
    }
}
